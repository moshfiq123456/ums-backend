package user_permissions

import (
	"context"

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

func (r *Repository) AssignPermissions(ctx context.Context, userID uuid.UUID, permissionIDs []uint) error {
	for _, pid := range permissionIDs {
		if err := r.db.WithContext(ctx).FirstOrCreate(&models.UserPermission{}, models.UserPermission{
			UserID:       userID,
			PermissionID: pid,
			Allow:        true,
		}).Error; err != nil {
			return err
		}
	}
	return nil
}

func (r *Repository) RemovePermissions(ctx context.Context, userID uuid.UUID, permissionIDs []uint) error {
	return r.db.WithContext(ctx).
		Where("user_id = ? AND permission_id IN ?", userID, permissionIDs).
		Delete(&models.UserPermission{}).Error
}

func (r *Repository) ListPermissions(ctx context.Context, userID uuid.UUID) ([]models.Permission, error) {
	var perms []models.Permission
	err := r.db.WithContext(ctx).
		Model(&models.Permission{}).
		Joins("JOIN user_permissions up ON up.permission_id = permissions.id").
		Where("up.user_id = ?", userID).
		Find(&perms).Error
	return perms, err
}
