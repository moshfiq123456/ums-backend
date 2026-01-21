package permissions

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

func (r *Repository) Create(ctx context.Context, p models.Permission) error {
	return r.db.WithContext(ctx).Create(&p).Error
}

func (r *Repository) List(ctx context.Context) ([]models.Permission, error) {
	var perms []models.Permission
	err := r.db.WithContext(ctx).Find(&perms).Error
	return perms, err
}

func (r *Repository) GetByID(ctx context.Context, id uint) (models.Permission, error) {
	var p models.Permission
	err := r.db.WithContext(ctx).First(&p, id).Error
	return p, err
}

func (r *Repository) Update(ctx context.Context, p models.Permission) error {
	return r.db.WithContext(ctx).Save(&p).Error
}

func (r *Repository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.Permission{}, id).Error
}
