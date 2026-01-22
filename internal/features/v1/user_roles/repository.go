package user_roles

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

func (r *Repository) AssignRoles(ctx context.Context, userID uuid.UUID, roleIDs []uint) error {
	for _, roleID := range roleIDs {
		if err := r.db.WithContext(ctx).FirstOrCreate(&models.UserRole{}, models.UserRole{
			UserID: userID,
			RoleID: roleID,
		}).Error; err != nil {
			return err
		}
	}
	return nil
}

func (r *Repository) RemoveRoles(ctx context.Context, userID uuid.UUID, roleIDs []uint) error {
	return r.db.WithContext(ctx).
		Where("user_id = ? AND role_id IN ?", userID, roleIDs).
		Delete(&models.UserRole{}).Error
}

func (r *Repository) ListRoles(
	ctx context.Context,
	userID uuid.UUID,
	page, size int,
) ([]models.Role, error) {

	var roles []models.Role
	offset := (page - 1) * size

	err := r.db.WithContext(ctx).
		Model(&models.Role{}).
		Joins("JOIN user_roles ur ON ur.role_id = roles.id").
		Where("ur.user_id = ?", userID).
		Order("roles.created_at DESC").
		Limit(size).
		Offset(offset).
		Find(&roles).Error

	return roles, err
}

