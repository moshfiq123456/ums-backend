// internal/features/v1/roles/repository.go
package roles

import (
	"context"

	"github.com/moshfiq123456/ums-backend/internal/models"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(ctx context.Context, role models.Role) (models.Role, error) {
	err := r.db.WithContext(ctx).Create(&role).Error
	return role, err
}

func (r *Repository) List(ctx context.Context, page, size int) ([]models.Role, error) {
	var roles []models.Role

	offset := (page - 1) * size

	err := r.db.WithContext(ctx).
		Order("created_at DESC").
		Limit(size).
		Offset(offset).
		Find(&roles).Error

	return roles, err
}


func (r *Repository) GetByID(ctx context.Context, id int64) (models.Role, error) {
	var role models.Role
	err := r.db.WithContext(ctx).First(&role, id).Error
	return role, err
}

func (r *Repository) Update(ctx context.Context, role models.Role) (models.Role, error) {
	err := r.db.WithContext(ctx).Save(&role).Error
	return role, err
}

func (r *Repository) UpdateStatus(ctx context.Context, id int64, isActive bool) error {
	return r.db.WithContext(ctx).Model(&models.Role{}).Where("id = ?", id).Update("is_active", isActive).Error
}

func (r *Repository) Delete(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Delete(&models.Role{}, id).Error
}
