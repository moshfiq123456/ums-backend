package sessions

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/moshfiq123456/ums-backend/internal/models"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) ListActive(ctx context.Context, page, size int, f ActiveSessionFilter) ([]models.LoginSession, int64, error) {
	var sessions []models.LoginSession
	var total int64
	offset := (page - 1) * size

	q := r.db.WithContext(ctx).
		Model(&models.LoginSession{}).
		Joins("JOIN users u ON u.id = login_sessions.user_id").
		Joins("JOIN organizations o ON o.id = u.organization_id").
		Where("login_sessions.logged_out_at IS NULL").
		Where("login_sessions.refresh_expires_at > ?", time.Now())

	if f.OrgID != "" {
		q = q.Where("u.organization_id = ?", f.OrgID)
	}
	if f.Search != "" {
		like := "%" + f.Search + "%"
		q = q.Where("u.name ILIKE ? OR u.email ILIKE ?", like, like)
	}

	q.Count(&total)

	err := q.Preload("User.Organization").
		Order("login_sessions.logged_in_at DESC").
		Limit(size).Offset(offset).
		Find(&sessions).Error

	return sessions, total, err
}

func (r *Repository) ForceLogout(ctx context.Context, sessionID uuid.UUID, orgID uuid.UUID) error {
	now := time.Now()
	return r.db.WithContext(ctx).
		Model(&models.LoginSession{}).
		Joins("JOIN users u ON u.id = login_sessions.user_id").
		Where("login_sessions.id = ? AND u.organization_id = ?", sessionID, orgID).
		Update("logged_out_at", now).Error
}
