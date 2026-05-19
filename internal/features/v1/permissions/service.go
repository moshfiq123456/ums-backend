package permissions

import (
	"context"

	"github.com/google/uuid"
	"github.com/moshfiq123456/ums-backend/internal/models"
	"github.com/moshfiq123456/ums-backend/internal/utils"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(ctx context.Context, req CreatePermissionRequest, orgID uuid.UUID) error {
	p := models.Permission{
		OrganizationID: orgID,
		Code:           req.Code,
		Name:           req.Name,
		Description:    req.Description,
	}
	return s.repo.Create(ctx, p)
}

func (s *Service) List(ctx context.Context, p utils.Pagination, f PermissionFilter) ([]models.Permission, int64, error) {
	return s.repo.List(ctx, p.Page, p.Size, f)
}

func (s *Service) Get(ctx context.Context, id uint, orgID uuid.UUID) (models.Permission, error) {
	return s.repo.GetByID(ctx, id, orgID)
}

func (s *Service) Update(ctx context.Context, id uint, req UpdatePermissionRequest, orgID uuid.UUID) (models.Permission, error) {
	p, err := s.repo.GetByID(ctx, id, orgID)
	if err != nil {
		return p, err
	}

	p.Name = req.Name
	p.Description = req.Description

	return p, s.repo.Update(ctx, p)
}

func (s *Service) Delete(ctx context.Context, id uint, orgID uuid.UUID) error {
	return s.repo.Delete(ctx, id, orgID)
}
