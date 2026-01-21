package users

import (
	"context"
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
func (r *UserRepository) List(ctx context.Context) ([]models.User, error) {
	var users []models.User
	err := r.db.WithContext(ctx).
		Where("deleted_at IS NULL").
		Order("created_at DESC").
		Find(&users).Error
	return users, err
}

// GET BY ID
func (r *UserRepository) GetByID(ctx context.Context, id string) (models.User, error) {
	var user models.User
	err := r.db.WithContext(ctx).
		Where("id = ? AND deleted_at IS NULL", id).
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
