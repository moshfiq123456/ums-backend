package roles

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/moshfiq123456/ums-backend/internal/middleware"
	"github.com/moshfiq123456/ums-backend/internal/utils"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

// POST /roles
func (h *Handler) Create(c *gin.Context) {
	var req CreateRoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	orgID, _ := middleware.GetOrgID(c)
	role, err := h.service.Create(c.Request.Context(), req, orgID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, toResponse(role))
}

// GET /roles
func (h *Handler) List(c *gin.Context) {
	var pagination utils.Pagination
	var filter RoleFilter
	_ = c.ShouldBindQuery(&pagination)
	_ = c.ShouldBindQuery(&filter)

	if orgID, ok := middleware.GetOrgID(c); ok {
		filter.OrgID = orgID.String()
	}

	if pagination.Page < 0 || pagination.Size < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid pagination params"})
		return
	}
	pagination.Normalize()

	roles, total, err := h.service.List(c.Request.Context(), pagination, filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": toResponseList(roles),
		"meta": gin.H{"page": pagination.Page, "size": pagination.Size, "total": total},
	})
}

// GET /roles/:id
func (h *Handler) Get(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("roleId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid role id"})
		return
	}

	role, err := h.service.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, toResponse(role))
}

// PUT /roles/:id
func (h *Handler) Update(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("roleId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid role id"})
		return
	}

	var req UpdateRoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	role, err := h.service.Update(c.Request.Context(), id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, toResponse(role))
}

// PATCH /roles/:id/status
func (h *Handler) SetStatus(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("roleId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid role id"})
		return
	}

	var req UpdateRoleStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.UpdateStatus(
		c.Request.Context(),
		id,
		req.IsActive,
	); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Role status updated successfully",
	})
}

// DELETE /roles/:id
func (h *Handler) Delete(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("roleId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid role id"})
		return
	}

	if err := h.service.Delete(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Role deleted successfully",
	})
}
