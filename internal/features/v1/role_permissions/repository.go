package role_permissions

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

func (r *Repository) Exists(ctx context.Context, roleID, permissionID uint) bool {
	var count int64
	r.db.WithContext(ctx).
		Model(&models.RolePermission{}).
		Where("role_id = ? AND permission_id = ?", roleID, permissionID).
		Count(&count)
	return count > 0
}

func (r *Repository) Create(ctx context.Context, roleID, permissionID uint) error {
	return r.db.WithContext(ctx).Create(&models.RolePermission{
		RoleID:       roleID,
		PermissionID: permissionID,
	}).Error
}

func (r *Repository) Delete(ctx context.Context, roleID, permissionID uint) error {
	return r.db.WithContext(ctx).
		Where("role_id = ? AND permission_id = ?", roleID, permissionID).
		Delete(&models.RolePermission{}).Error
}

func (r *Repository) List(
	ctx context.Context,
	roleID uint,
	page, size int,
) ([]models.Permission, error) {

	var perms []models.Permission
	offset := (page - 1) * size

	err := r.db.WithContext(ctx).
		Where("role_id = ?", roleID).
		Order("created_at DESC").
		Limit(size).
		Offset(offset).
		Find(&perms).Error

	return perms, err
}

