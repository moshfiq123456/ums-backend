package role_permissions

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/moshfiq123456/ums-backend/internal/utils"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

// POST /roles/:roleId/permissions
func (h *Handler) Assign(c *gin.Context) {
	roleID, err := strconv.ParseUint(c.Param("roleId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid role id"})
		return
	}

	var req AssignPermissionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.Assign(c.Request.Context(), uint(roleID), req.PermissionID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Permission assigned to role successfully",
	})
}

// POST /roles/:roleId/permissions/bulk
func (h *Handler) BulkAssign(c *gin.Context) {
	roleID, err := strconv.ParseUint(c.Param("roleId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid role id"})
		return
	}

	var req BulkAssignPermissionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.BulkAssign(
		c.Request.Context(),
		uint(roleID),
		req.PermissionIDs,
	); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Permissions assigned to role successfully",
	})
}

// DELETE /roles/:roleId/permissions/:permissionId
func (h *Handler) Remove(c *gin.Context) {
	roleID, err := strconv.ParseUint(c.Param("roleId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid role id"})
		return
	}

	permissionID, err := strconv.ParseUint(c.Param("permissionId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid permission id"})
		return
	}

	if err := h.service.Remove(
		c.Request.Context(),
		uint(roleID),
		uint(permissionID),
	); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Permission removed from role successfully",
	})
}

// GET /roles/:roleId/permissions
func (h *Handler) List(c *gin.Context) {
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

	// 4️⃣ Parse path param
	roleID, err := strconv.ParseUint(c.Param("roleId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid role id",
		})
		return
	}

	// 5️⃣ Call service
	result, err := h.service.List(
		c.Request.Context(),
		uint(roleID),
		pagination,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 6️⃣ Response
	c.JSON(http.StatusOK, gin.H{
		"data": result,
		"meta": gin.H{
			"page": pagination.Page,
			"size": pagination.Size,
		},
	})
}

