package auth

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

func (r *Repository) FindUserByEmail(ctx context.Context, email string) (models.User, error) {
	var user models.User
	err := r.db.WithContext(ctx).
		Where("email = ?", email).
		First(&user).Error
	return user, err
}

func (r *Repository) CreateSession(ctx context.Context, session models.LoginSession) error {
	return r.db.WithContext(ctx).Create(&session).Error
}

func (r *Repository) FindSessionByID(ctx context.Context, id uuid.UUID) (models.LoginSession, error) {
	var session models.LoginSession
	err := r.db.WithContext(ctx).
		Where("id = ?", id).
		Where("logged_out_at IS NULL").
		First(&session).Error
	return session, err
}

func (r *Repository) LogoutSession(ctx context.Context, id uuid.UUID) error {
	now := time.Now()
	return r.db.WithContext(ctx).
		Model(&models.LoginSession{}).
		Where("id = ?", id).
		Update("logged_out_at", now).Error
}
