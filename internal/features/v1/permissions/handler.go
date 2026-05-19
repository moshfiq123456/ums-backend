package permissions

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/moshfiq123456/ums-backend/internal/constants"
	"github.com/moshfiq123456/ums-backend/internal/middleware"
	"github.com/moshfiq123456/ums-backend/internal/utils"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

// CreatePermission godoc
func (h *Handler) Create(c *gin.Context) {
	var req CreatePermissionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Priority: body org_id → JWT org_id → default org
	orgID := constants.DefaultOrgID
	if req.OrgID != "" {
		orgID = uuid.MustParse(req.OrgID)
	} else if jwtOrgID, ok := middleware.GetOrgID(c); ok {
		orgID = jwtOrgID
	}

	if err := h.service.Create(c.Request.Context(), req, orgID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "permission created"})
}

// ListPermissions godoc
func (h *Handler) List(c *gin.Context) {
	var pagination utils.Pagination
	var filter PermissionFilter
	_ = c.ShouldBindQuery(&pagination)
	_ = c.ShouldBindQuery(&filter)

	// Priority: query param org_id → JWT org_id
	if filter.OrgID == "" {
		if orgID, ok := middleware.GetOrgID(c); ok {
			filter.OrgID = orgID.String()
		}
	}

	if pagination.Page < 0 || pagination.Size < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid pagination params"})
		return
	}
	pagination.Normalize()

	perms, total, err := h.service.List(c.Request.Context(), pagination, filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": ToResponseList(perms),
		"meta": gin.H{"page": pagination.Page, "size": pagination.Size, "total": total},
	})
}

// GetPermission godoc
func (h *Handler) Get(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid permission id"})
		return
	}

	orgID, ok := middleware.GetOrgID(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "organization context required"})
		return
	}

	p, err := h.service.Get(c.Request.Context(), uint(id), orgID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, ToResponse(p))
}

// UpdatePermission godoc
func (h *Handler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid permission id"})
		return
	}

	orgID, ok := middleware.GetOrgID(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "organization context required"})
		return
	}

	var req UpdatePermissionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.service.Update(c.Request.Context(), uint(id), req, orgID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, ToResponse(resp))
}

// DeletePermission godoc
func (h *Handler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid permission id"})
		return
	}

	orgID, ok := middleware.GetOrgID(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "organization context required"})
		return
	}

	if err := h.service.Delete(c.Request.Context(), uint(id), orgID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "permission deleted"})
}
