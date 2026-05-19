package users

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

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

// CreateUser godoc
func (h *Handler) CreateUser(c *gin.Context) {
	var req CreateUserRequest
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

	user, err := h.service.Create(c.Request.Context(), req, orgID)
	if err != nil {
		if err.Error() == "validation failed" {
			c.JSON(http.StatusBadRequest, gin.H{"error": utils.FormatValidationError(err)})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, toResponse(user))
}

// UpdateUser godoc
func (h *Handler) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var req UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.service.Update(c.Request.Context(), id, req)
	if err != nil {
		if err.Error() == "validation failed" {
			c.JSON(http.StatusBadRequest, gin.H{"error": utils.FormatValidationError(err)})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, toResponse(user))
}



// ListUsers godoc
func (h *Handler) ListUsers(c *gin.Context) {
	var pagination utils.Pagination
	var filter UserFilter
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

	users, total, err := h.service.List(c.Request.Context(), pagination, filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": toResponseList(users),
		"meta": gin.H{"page": pagination.Page, "size": pagination.Size, "total": total},
	})
}


// GetUser godoc
func (h *Handler) GetUser(c *gin.Context) {
	id := c.Param("id")
	orgID, ok := middleware.GetOrgID(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "organization context required"})
		return
	}
	user, err := h.service.GetByID(c.Request.Context(), id, orgID.String())
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, toResponse(user))
}

// DeleteUser godoc
func (h *Handler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.Delete(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "user deleted"})
}

// SetStatus godoc
func (h *Handler) SetStatus(c *gin.Context) {
	id := c.Param("id")
	var req UpdateStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.UpdateStatus(c.Request.Context(), id, req.Status); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "status updated"})
}

// ChangePassword godoc
func (h *Handler) ChangePassword(c *gin.Context) {
	id := c.Param("id")
	var req ChangePasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.ChangePassword(c.Request.Context(), id, req.Password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "password updated"})
}

// PATCH /users/:id/avatar
func (h *Handler) UploadAvatar(c *gin.Context) {
	id := c.Param("id")

	file, header, err := c.Request.FormFile("avatar")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "avatar file is required"})
		return
	}
	defer file.Close()

	// Validate size (2 MB)
	if header.Size > 2<<20 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "file too large, max 2 MB"})
		return
	}

	// Validate MIME type
	ext := strings.ToLower(filepath.Ext(header.Filename))
	allowed := map[string]bool{".jpg": true, ".jpeg": true, ".png": true, ".webp": true}
	if !allowed[ext] {
		c.JSON(http.StatusBadRequest, gin.H{"error": "only JPEG, PNG and WebP are allowed"})
		return
	}

	orgID, ok := middleware.GetOrgID(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "organization context required"})
		return
	}

	// Fetch current user to get old avatar path
	existing, _ := h.service.GetByID(c.Request.Context(), id, orgID.String())

	// Ensure upload directory exists
	uploadDir := "./uploads/avatars"
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create upload directory"})
		return
	}

	// Save new file with UUID filename
	filename := fmt.Sprintf("%s%s", uuid.New().String(), ext)
	destPath := filepath.Join(uploadDir, filename)
	if err := c.SaveUploadedFile(header, destPath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not save file"})
		return
	}

	avatarURL := fmt.Sprintf("/uploads/avatars/%s", filename)

	if err := h.service.UpdateAvatar(c.Request.Context(), id, avatarURL); err != nil {
		_ = os.Remove(destPath) // rollback new file if DB update fails
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Delete old file after successful DB update
	if existing.AvatarURL != nil && *existing.AvatarURL != "" {
		oldPath := "." + *existing.AvatarURL // "/uploads/avatars/xxx.jpg" → "./uploads/avatars/xxx.jpg"
		_ = os.Remove(oldPath)
	}

	c.JSON(http.StatusOK, gin.H{"avatar_url": avatarURL})
}

// DELETE /users/:id/avatar
func (h *Handler) RemoveAvatar(c *gin.Context) {
	id := c.Param("id")

	orgID, ok := middleware.GetOrgID(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "organization context required"})
		return
	}

	// Fetch user to get the old file path
	user, err := h.service.GetByID(c.Request.Context(), id, orgID.String())
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	if user.AvatarURL == nil || *user.AvatarURL == "" {
		c.JSON(http.StatusOK, gin.H{"message": "no avatar to remove"})
		return
	}

	oldPath := "." + *user.AvatarURL
	if err := h.service.UpdateAvatar(c.Request.Context(), id, ""); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	_ = os.Remove(oldPath)
	c.JSON(http.StatusOK, gin.H{"message": "avatar removed"})
}
