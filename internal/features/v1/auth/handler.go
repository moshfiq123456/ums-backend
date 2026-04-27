package auth

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func setRefreshCookie(c *gin.Context, value string, maxAge int) {
	c.SetSameSite(http.SameSiteNoneMode)
	c.SetCookie(refreshTokenCookie, value, maxAge, "/", "", false, true)
}

func clearRefreshCookie(c *gin.Context) {
	c.SetSameSite(http.SameSiteNoneMode)
	c.SetCookie(refreshTokenCookie, "", -1, "/", "", false, true)
}

const refreshTokenCookie = "refresh_token"

type Handler struct {
	service *Service
}

func NewHandler(s *Service) *Handler {
	return &Handler{service: s}
}

// --------------------
// LOGIN
// --------------------
func (h *Handler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.service.Login(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Set refresh token as httpOnly cookie
	maxAge := int(parseDuration(os.Getenv("REFRESH_TOKEN_TTL")).Seconds())
	if maxAge <= 0 {
		maxAge = int((7 * 24 * time.Hour).Seconds())
	}
	setRefreshCookie(c, resp.RefreshToken, maxAge)

	// Don't expose refresh token in body
	resp.RefreshToken = ""
	c.JSON(http.StatusOK, resp)
}

// --------------------
// REFRESH
// --------------------
func (h *Handler) Refresh(c *gin.Context) {
	cookie, err := c.Cookie(refreshTokenCookie)
	if err != nil || cookie == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "missing refresh token"})
		return
	}

	resp, err := h.service.Refresh(c.Request.Context(), RefreshTokenRequest{RefreshToken: cookie})
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Rotate cookie
	maxAge := int(parseDuration(os.Getenv("REFRESH_TOKEN_TTL")).Seconds())
	if maxAge <= 0 {
		maxAge = int((7 * 24 * time.Hour).Seconds())
	}
	setRefreshCookie(c, resp.RefreshToken, maxAge)

	resp.RefreshToken = ""
	c.JSON(http.StatusOK, gin.H{"access_token": resp.AccessToken})
}

// --------------------
// LOGOUT
// --------------------
func (h *Handler) Logout(c *gin.Context) {
	cookie, err := c.Cookie(refreshTokenCookie)
	if err != nil || cookie == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "missing refresh token"})
		return
	}

	if err := h.service.Logout(c.Request.Context(), LogoutRequest{RefreshToken: cookie}); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Clear cookie
	clearRefreshCookie(c)
	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}
