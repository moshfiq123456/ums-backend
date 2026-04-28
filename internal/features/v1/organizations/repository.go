package organizations

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

func (r *Repository) Create(ctx context.Context, org models.Organization) (models.Organization, error) {
	err := r.db.WithContext(ctx).Create(&org).Error
	return org, err
}

func (r *Repository) List(ctx context.Context, page, size int, f OrgFilter) ([]models.Organization, int64, error) {
	var orgs []models.Organization
	var total int64

	q := r.db.WithContext(ctx).Model(&models.Organization{}).Where("deleted_at IS NULL")
	if f.Search != "" {
		like := "%" + f.Search + "%"
		q = q.Where("name ILIKE ? OR slug ILIKE ? OR domain ILIKE ?", like, like, like)
	}
	if f.IsActive != nil {
		q = q.Where("is_active = ?", *f.IsActive)
	}
	if f.Plan != "" {
		q = q.Where("plan = ?", f.Plan)
	}
	q.Count(&total)

	offset := (page - 1) * size
	err := q.Order("created_at DESC").Limit(size).Offset(offset).Find(&orgs).Error
	return orgs, total, err
}

func (r *Repository) GetByID(ctx context.Context, id uuid.UUID) (models.Organization, error) {
	var org models.Organization
	err := r.db.WithContext(ctx).Where("id = ? AND deleted_at IS NULL", id).First(&org).Error
	return org, err
}

func (r *Repository) GetBySlug(ctx context.Context, slug string) (models.Organization, error) {
	var org models.Organization
	err := r.db.WithContext(ctx).Where("slug = ? AND deleted_at IS NULL", slug).First(&org).Error
	return org, err
}

func (r *Repository) Update(ctx context.Context, org models.Organization) (models.Organization, error) {
	err := r.db.WithContext(ctx).Save(&org).Error
	return org, err
}

func (r *Repository) UpdateStatus(ctx context.Context, id uuid.UUID, isActive bool) error {
	return r.db.WithContext(ctx).Model(&models.Organization{}).
		Where("id = ?", id).Update("is_active", isActive).Error
}

func (r *Repository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Model(&models.Organization{}).
		Where("id = ?", id).Update("deleted_at", gorm.Expr("NOW()")).Error
}

func (r *Repository) CountUsers(ctx context.Context, id uuid.UUID) int64 {
	var count int64
	r.db.WithContext(ctx).Table("users").Where("organization_id = ? AND deleted_at IS NULL", id).Count(&count)
	return count
}

func (r *Repository) CountRoles(ctx context.Context, id uuid.UUID) int64 {
	var count int64
	r.db.WithContext(ctx).Table("roles").Where("organization_id = ?", id).Count(&count)
	return count
}

func (r *Repository) SlugExists(ctx context.Context, slug string, excludeID *uuid.UUID) bool {
	q := r.db.WithContext(ctx).Table("organizations").Where("slug = ? AND deleted_at IS NULL", strings.ToLower(slug))
	if excludeID != nil {
		q = q.Where("id != ?", *excludeID)
	}
	var count int64
	q.Count(&count)
	return count > 0
}
