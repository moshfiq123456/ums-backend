package user_hierarchy

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

// POST /users/:parentId/children
func (h *Handler) AssignChild(c *gin.Context) {
	parentID, err := uuid.Parse(c.Param("parentId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid parent id"})
		return
	}

	var req AssignChildRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.AssignChild(
		c.Request.Context(),
		parentID,
		req.ChildUserID,
	); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Hierarchy created successfully",
	})
}

// DELETE /users/:parentId/children/:childId
func (h *Handler) RemoveChild(c *gin.Context) {
	parentID, err := uuid.Parse(c.Param("parentId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid parent id"})
		return
	}

	childID, err := uuid.Parse(c.Param("childId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid child id"})
		return
	}

	if err := h.service.RemoveChild(
		c.Request.Context(),
		parentID,
		childID,
	); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Hierarchy removed successfully",
	})
}

// GET /users/:id/children
func (h *Handler) GetChildren(c *gin.Context) {
	userID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	var pagination utils.Pagination
	_ = c.ShouldBindQuery(&pagination)

	if pagination.Page < 0 || pagination.Size < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid pagination params"})
		return
	}

	pagination.Normalize()

	children, err := h.service.GetChildren(c.Request.Context(), userID, pagination)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": children,
		"meta": gin.H{
			"page": pagination.Page,
			"size": pagination.Size,
		},
	})
}

// GET /users/:id/parent
func (h *Handler) GetParent(c *gin.Context) {
	userID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	parent, err := h.service.GetParent(
		c.Request.Context(),
		userID,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, parent)
}

// POST /users/hierarchy/check
func (h *Handler) CheckHierarchy(c *gin.Context) {
	var req CheckHierarchyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.service.CheckHierarchy(
		c.Request.Context(),
		req.ParentUserID,
		req.ChildUserID,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}
