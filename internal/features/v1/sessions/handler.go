package sessions

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

// ListActive godoc
func (h *Handler) ListActive(c *gin.Context) {
	var p utils.Pagination
	var f ActiveSessionFilter
	_ = c.ShouldBindQuery(&p)
	_ = c.ShouldBindQuery(&f)

	if f.OrgID == "" {
		if orgID, ok := middleware.GetOrgID(c); ok {
			f.OrgID = orgID.String()
		}
	}
	p.Normalize()

	sessions, total, err := h.service.ListActive(c.Request.Context(), p, f)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": sessions,
		"meta": gin.H{"page": p.Page, "size": p.Size, "total": total},
	})
}

// ForceLogout godoc
func (h *Handler) ForceLogout(c *gin.Context) {
	sessionID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid session id"})
		return
	}

	orgID, ok := middleware.GetOrgID(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "organization context required"})
		return
	}

	if err := h.service.ForceLogout(c.Request.Context(), sessionID, orgID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "session terminated"})
}
