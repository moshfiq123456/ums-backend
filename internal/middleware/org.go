package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

const OrgIDKey = "org_id"

// OrgContext extracts organization_id from the JWT and sets it in the Gin context.
// Must run after JWTAuth.
func OrgContext() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.Next()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 {
			c.Next()
			return
		}

		token, err := jwt.Parse(parts[1], func(t *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("ACCESS_TOKEN_SECRET")), nil
		})
		if err != nil || !token.Valid {
			c.Next()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.Next()
			return
		}

		if orgIDStr, ok := claims["org_id"].(string); ok {
			if orgID, err := uuid.Parse(orgIDStr); err == nil {
				c.Set(OrgIDKey, orgID)
			}
		}
		c.Next()
	}
}

// RequireOrg aborts the request if organization_id is not in context.
func RequireOrg() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, exists := c.Get(OrgIDKey); !exists {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "organization context required"})
			return
		}
		c.Next()
	}
}

// GetOrgID retrieves the organization UUID from Gin context.
func GetOrgID(c *gin.Context) (uuid.UUID, bool) {
	val, exists := c.Get(OrgIDKey)
	if !exists {
		return uuid.Nil, false
	}
	orgID, ok := val.(uuid.UUID)
	return orgID, ok
}
