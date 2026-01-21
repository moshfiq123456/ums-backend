package user_hierarchy

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
		Joins("JOIN user_hierarchies uh ON uh.child_user_id = users.id").
		Where("uh.parent_user_id = ?", userID).
		Find(&users).Error
	return users, err
}

func (r *Repository) GetParent(ctx context.Context, userID uuid.UUID) (models.User, error) {
	var user models.User
	err := r.db.WithContext(ctx).
		Joins("JOIN user_hierarchies uh ON uh.parent_user_id = users.id").
		Where("uh.child_user_id = ?", userID).
		First(&user).Error
	return user, err
}
