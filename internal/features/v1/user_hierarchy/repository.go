package user_hierarchy

import (
	"context"
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

func (r *Repository) Exists(ctx context.Context, parentID, childID uuid.UUID) bool {
	var count int64
	r.db.WithContext(ctx).
		Model(&models.UserHierarchy{}).
		Where("parent_user_id = ? AND child_user_id = ?", parentID, childID).
		Count(&count)
	return count > 0
}

func (r *Repository) Create(ctx context.Context, parentID, childID uuid.UUID) error {
	return r.db.WithContext(ctx).Create(&models.UserHierarchy{
		ParentUserID: parentID,
		ChildUserID:  childID,
	}).Error
}

func (r *Repository) Delete(ctx context.Context, parentID, childID uuid.UUID) error {
	return r.db.WithContext(ctx).
		Where("parent_user_id = ? AND child_user_id = ?", parentID, childID).
		Delete(&models.UserHierarchy{}).Error
}

func (r *Repository) GetChildren(ctx context.Context, userID uuid.UUID) ([]models.User, error) {
	var users []models.User
	err := r.db.WithContext(ctx).
		Joins("JOIN user_hierarchy uh ON uh.child_user_id = users.id").
		Where("uh.parent_user_id = ?", userID).
		Find(&users).Error
	return users, err
}

func (r *Repository) GetParent(ctx context.Context, userID uuid.UUID) (models.User, error) {
	var user models.User
	err := r.db.WithContext(ctx).
		Joins("JOIN user_hierarchy uh ON uh.parent_user_id = users.id").
		Where("uh.child_user_id = ?", userID).
		First(&user).Error
	return user, err
}

func (r *Repository) ListAll(ctx context.Context, page, size int, f HierarchyFilter) ([]HierarchyDetail, int64, error) {
	var total int64
	base := r.db.WithContext(ctx).
		Table("user_hierarchy uh").
		Joins("JOIN users p ON p.id = uh.parent_user_id").
		Joins("JOIN users c ON c.id = uh.child_user_id").
		Where("p.deleted_at IS NULL AND c.deleted_at IS NULL")

	if f.OrgID != "" {
		base = base.Where("p.organization_id = ? AND c.organization_id = ?", f.OrgID, f.OrgID)
	}
	if f.Search != "" {
		like := "%" + f.Search + "%"
		base = base.Where("p.name ILIKE ? OR p.email ILIKE ? OR c.name ILIKE ? OR c.email ILIKE ?", like, like, like, like)
	}
	if f.ParentID != "" {
		base = base.Where("p.id IN ?", strings.Split(f.ParentID, ","))
	}
	if f.ChildID != "" {
		base = base.Where("c.id IN ?", strings.Split(f.ChildID, ","))
	}

	base.Count(&total)

	var results []HierarchyDetail
	offset := (page - 1) * size
	err := base.
		Select("p.id as parent_id, p.name as parent_name, p.email as parent_email, c.id as child_id, c.name as child_name, c.email as child_email").
		Order("p.name ASC, c.name ASC").
		Limit(size).Offset(offset).
		Scan(&results).Error

	return results, total, err
}
