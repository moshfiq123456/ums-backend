package user_roles

import (
	"context"
	"errors"
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

func (r *Repository) AssignRoles(ctx context.Context, userID uuid.UUID, orgID uuid.UUID, roleIDs []uint) error {
	var userCount int64
	r.db.WithContext(ctx).Model(&models.User{}).
		Where("id = ? AND organization_id = ? AND deleted_at IS NULL", userID, orgID).
		Count(&userCount)
	if userCount == 0 {
		return errors.New("user not found in organization")
	}

	var validCount int64
	r.db.WithContext(ctx).Model(&models.Role{}).
		Where("id IN ? AND organization_id = ?", roleIDs, orgID).
		Count(&validCount)
	if int(validCount) != len(roleIDs) {
		return errors.New("one or more roles not found in organization")
	}

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

func (r *Repository) RemoveRoles(ctx context.Context, userID uuid.UUID, orgID uuid.UUID, roleIDs []uint) error {
	return r.db.WithContext(ctx).
		Where("user_id = ? AND role_id IN ? AND EXISTS (SELECT 1 FROM users WHERE id = ? AND organization_id = ?)", userID, roleIDs, userID, orgID).
		Delete(&models.UserRole{}).Error
}

func (r *Repository) ListRoles(ctx context.Context, userID uuid.UUID, orgID uuid.UUID, page, size int) ([]models.Role, error) {
	var roles []models.Role
	offset := (page - 1) * size

	err := r.db.WithContext(ctx).
		Model(&models.Role{}).
		Joins("JOIN user_roles ur ON ur.role_id = roles.id").
		Joins("JOIN users u ON u.id = ur.user_id").
		Where("ur.user_id = ? AND u.organization_id = ?", userID, orgID).
		Order("roles.created_at DESC").
		Limit(size).Offset(offset).
		Find(&roles).Error

	return roles, err
}

func (r *Repository) ListAll(ctx context.Context, page, size int, f UserRoleFilter) ([]UserRoleDetail, int64, error) {
	var total int64
	base := r.db.WithContext(ctx).
		Table("user_roles ur").
		Joins("JOIN users u ON u.id = ur.user_id").
		Joins("JOIN roles  r ON r.id = ur.role_id").
		Where("u.deleted_at IS NULL")

	if f.OrgID != "" {
		base = base.Where("u.organization_id = ?", f.OrgID)
	}

	if f.Search != "" {
		like := "%" + f.Search + "%"
		base = base.Where("u.name ILIKE ? OR u.email ILIKE ? OR r.name ILIKE ? OR r.code ILIKE ?", like, like, like, like)
	}
	if f.UserID != "" {
		base = base.Where("u.id IN ?", strings.Split(f.UserID, ","))
	}
	if f.RoleID != "" {
		var ids []uint
		for _, s := range strings.Split(f.RoleID, ",") {
			s = strings.TrimSpace(s)
			if s != "" {
				var id uint
				if _, err := fmt.Sscanf(s, "%d", &id); err == nil {
					ids = append(ids, id)
				}
			}
		}
		if len(ids) > 0 {
			base = base.Where("r.id IN ?", ids)
		}
	}

	base.Count(&total)

	var results []UserRoleDetail
	offset := (page - 1) * size
	err := base.
		Joins("JOIN organizations o ON o.id = u.organization_id").
		Select("o.id as org_id, o.name as org_name, u.id as user_id, u.name as user_name, u.email as user_email, u.avatar_url as user_avatar_url, r.id as role_id, r.name as role_name, r.code as role_code").
		Order("u.name ASC, r.name ASC").
		Limit(size).Offset(offset).
		Scan(&results).Error

	return results, total, err
}

