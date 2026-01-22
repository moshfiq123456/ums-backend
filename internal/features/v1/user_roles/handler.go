package user_roles

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/moshfiq123456/ums-backend/internal/utils"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

// POST /users/:id/roles
func (h *Handler) AssignRoles(c *gin.Context) {
	userID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	var req AssignRolesRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.AssignRoles(c.Request.Context(), userID, req.RoleIDs); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Roles assigned to user successfully"})
}

// DELETE /users/:id/roles
func (h *Handler) RemoveRoles(c *gin.Context) {
	userID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	var req RemoveRolesRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.RemoveRoles(c.Request.Context(), userID, req.RoleIDs); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Roles removed from user successfully"})
}

// GET /users/:id/roles
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
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid user id",
		})
		return
	}

	// 5️⃣ Call service
	roles, err := h.service.ListRoles(
		c.Request.Context(),
		userID,
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

