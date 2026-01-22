package user_permissions

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

// POST /users/:id/permissions
func (h *Handler) AssignPermissions(c *gin.Context) {
	userID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	var req AssignPermissionsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.AssignPermissions(
		c.Request.Context(),
		userID,
		req.PermissionIDs,
	); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Permissions assigned successfully",
	})
}

// DELETE /users/:id/permissions
func (h *Handler) RemovePermissions(c *gin.Context) {
	userID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	var req RemovePermissionsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.RemovePermissions(
		c.Request.Context(),
		userID,
		req.PermissionIDs,
	); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Permissions removed successfully",
	})
}

// GET /users/:id/permissions
func (h *Handler) ListPermissions(c *gin.Context) {
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
	perms, err := h.service.ListPermissions(
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
	resp := ToUserPermissionResponses(c.Param("id"), perms)

	// 7️⃣ Response with meta
	c.JSON(http.StatusOK, gin.H{
		"data": resp,
		"meta": gin.H{
			"page": pagination.Page,
			"size": pagination.Size,
		},
	})
}

