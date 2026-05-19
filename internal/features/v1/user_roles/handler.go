package user_roles

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/moshfiq123456/ums-backend/internal/middleware"
	"github.com/moshfiq123456/ums-backend/internal/utils"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

// AssignRoles godoc
func (h *Handler) AssignRoles(c *gin.Context) {
	userID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	orgID, ok := middleware.GetOrgID(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "organization context required"})
		return
	}

	var req AssignRolesRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.AssignRoles(c.Request.Context(), userID, orgID, req.RoleIDs); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Roles assigned to user successfully"})
}

// RemoveRoles godoc
func (h *Handler) RemoveRoles(c *gin.Context) {
	userID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	orgID, ok := middleware.GetOrgID(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "organization context required"})
		return
	}

	var req RemoveRolesRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.RemoveRoles(c.Request.Context(), userID, orgID, req.RoleIDs); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Roles removed from user successfully"})
}

// ListAllUserRoles godoc
func (h *Handler) ListAll(c *gin.Context) {
	var p utils.Pagination
	var f UserRoleFilter
	_ = c.ShouldBindQuery(&p)
	_ = c.ShouldBindQuery(&f)
	if f.OrgID == "" {
		if orgID, ok := middleware.GetOrgID(c); ok {
			f.OrgID = orgID.String()
		}
	}
	p.Normalize()

	results, total, err := h.service.ListAll(c.Request.Context(), p, f)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": results,
		"meta": gin.H{"page": p.Page, "size": p.Size, "total": total},
	})
}

// ListUserRoles godoc
func (h *Handler) ListRoles(c *gin.Context) {
	var pagination utils.Pagination

	// 1️⃣ Bind query params
	_ = c.ShouldBindQuery(&pagination)

	// 2️⃣ Hard validation
	if pagination.Page < 0 || pagination.Size < 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid pagination params",
		})
		return
	}

	// 3️⃣ Normalize defaults
	pagination.Normalize()

	// 4️⃣ Parse user ID
	userID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	orgID, ok := middleware.GetOrgID(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "organization context required"})
		return
	}

	// 5️⃣ Call service
	roles, err := h.service.ListRoles(
		c.Request.Context(),
		userID,
		orgID,
		pagination,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 6️⃣ Map response
	resp := ToUserRoleResponses(c.Param("id"), roles)

	// 7️⃣ Return with meta
	c.JSON(http.StatusOK, gin.H{
		"data": resp,
		"meta": gin.H{
			"page": pagination.Page,
			"size": pagination.Size,
		},
	})
}

