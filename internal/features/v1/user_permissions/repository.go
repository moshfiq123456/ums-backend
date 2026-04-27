package user_permissions

import (
	"context"
	"fmt"
	"strings"

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

func (r *Repository) ListPermissions(
	ctx context.Context,
	userID uuid.UUID,
	page, size int,
) ([]models.Permission, error) {

	var perms []models.Permission
	offset := (page - 1) * size

	err := r.db.WithContext(ctx).
		Model(&models.Permission{}).
		Joins("JOIN user_permissions up ON up.permission_id = permissions.id").
		Where("up.user_id = ?", userID).
		Order("permissions.created_at DESC").
		Limit(size).
		Offset(offset).
		Find(&perms).Error

	return perms, err
}

func (r *Repository) ListAll(ctx context.Context, page, size int, f UserPermissionFilter) ([]UserPermissionDetail, int64, error) {
	var total int64
	base := r.db.WithContext(ctx).
		Table("user_permissions up").
		Joins("JOIN users u ON u.id = up.user_id").
		Joins("JOIN permissions p ON p.id = up.permission_id").
		Where("u.deleted_at IS NULL")

	if f.Search != "" {
		like := "%" + f.Search + "%"
		base = base.Where("u.name ILIKE ? OR u.email ILIKE ? OR p.code ILIKE ? OR p.name ILIKE ?", like, like, like, like)
	}
	if f.UserID != "" {
		base = base.Where("u.id IN ?", strings.Split(f.UserID, ","))
	}
	if f.PermissionID != "" {
		var ids []uint
		for _, s := range strings.Split(f.PermissionID, ",") {
			var id uint
			if _, err := fmt.Sscanf(strings.TrimSpace(s), "%d", &id); err == nil {
				ids = append(ids, id)
			}
		}
		if len(ids) > 0 {
			base = base.Where("p.id IN ?", ids)
		}
	}

	base.Count(&total)

	var results []UserPermissionDetail
	offset := (page - 1) * size
	err := base.
		Select("u.id as user_id, u.name as user_name, u.email as user_email, u.avatar_url as user_avatar_url, p.id as permission_id, p.code, p.name, p.description").
		Order("u.name ASC, p.code ASC").
		Limit(size).Offset(offset).
		Scan(&results).Error

	return results, total, err
}

