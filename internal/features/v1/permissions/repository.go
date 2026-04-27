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

func (r *Repository) List(ctx context.Context, page, size int, f PermissionFilter) ([]models.Permission, int64, error) {
	var perms []models.Permission
	var total int64
	offset := (page - 1) * size

	q := r.db.WithContext(ctx).Model(&models.Permission{})
	if f.Search != "" {
		like := "%" + f.Search + "%"
		q = q.Where("name ILIKE ? OR code ILIKE ?", like, like)
	}
	if f.Service != "" {
		q = q.Where("service = ?", f.Service)
	}
	q.Count(&total)

	err := q.Order("created_at DESC").Limit(size).Offset(offset).Find(&perms).Error
	return perms, total, err
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
