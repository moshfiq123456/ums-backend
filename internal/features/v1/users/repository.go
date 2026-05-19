package users

import (
	"context"
	"strings"
	"time"

	"github.com/moshfiq123456/ums-backend/internal/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

// CREATE
func (r *UserRepository) Create(ctx context.Context, user models.User) (models.User, error) {
	err := r.db.WithContext(ctx).Create(&user).Error
	return user, err
}

// LIST
func (r *UserRepository) List(ctx context.Context, page, size int, f UserFilter) ([]models.User, int64, error) {
	var users []models.User
	var total int64
	offset := (page - 1) * size

	q := r.db.WithContext(ctx).Model(&models.User{}).Where("deleted_at IS NULL")
	// Scope by organization if provided
	if f.OrgID != "" {
		q = q.Where("organization_id = ?", f.OrgID)
	}
	if f.Search != "" {
		like := "%" + f.Search + "%"
		q = q.Where("name ILIKE ? OR email ILIKE ?", like, like)
	}
	if f.Status != "" {
		statuses := strings.Split(f.Status, ",")
		q = q.Where("status IN ?", statuses)
	}
	q.Count(&total)

	err := q.Preload("Roles").Preload("Organization").Order("created_at DESC").Limit(size).Offset(offset).Find(&users).Error
	return users, total, err
}

// GET BY ID
func (r *UserRepository) GetByID(ctx context.Context, id string, orgID string) (models.User, error) {
	var user models.User
	err := r.db.WithContext(ctx).
		Preload("Organization").
		Where("id = ? AND organization_id = ? AND deleted_at IS NULL", id, orgID).
		First(&user).Error
	return user, err
}

// UPDATE
func (r *UserRepository) Update(ctx context.Context, id string, name *string, phone *string) (models.User, error) {
	var user models.User
	if err := r.db.WithContext(ctx).First(&user, "id = ?", id).Error; err != nil {
		return user, err
	}

	if name != nil {
		user.Name = *name
	}
	if phone != nil {
		user.Phone = phone
	}

	err := r.db.WithContext(ctx).Save(&user).Error
	return user, err
}

// UPDATE STATUS
func (r *UserRepository) UpdateStatus(ctx context.Context, id string, status string) error {
	return r.db.WithContext(ctx).
		Model(&models.User{}).
		Where("id = ?", id).
		Update("status", status).Error
}

// SOFT DELETE
func (r *UserRepository) Delete(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).
		Model(&models.User{}).
		Where("id = ?", id).
		Update("deleted_at", time.Now()).Error
}

// GET BY EMAIL
func (r *UserRepository) GetByEmail(ctx context.Context, email string) (models.User, error) {
	var user models.User
	err := r.db.WithContext(ctx).Where("email = ? AND deleted_at IS NULL", email).First(&user).Error
	return user, err
}

// CHANGE PASSWORD
func (r *UserRepository) ChangePassword(ctx context.Context, id string, hash string) error {
	return r.db.WithContext(ctx).
		Model(&models.User{}).
		Where("id = ?", id).
		Update("password_hash", hash).Error
}

// UPDATE AVATAR (pass empty string to clear)
func (r *UserRepository) UpdateAvatar(ctx context.Context, id string, avatarURL string) error {
	val := interface{}(avatarURL)
	if avatarURL == "" {
		val = nil
	}
	return r.db.WithContext(ctx).
		Model(&models.User{}).
		Where("id = ?", id).
		Update("avatar_url", val).Error
}
