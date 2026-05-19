// internal/features/v1/roles/service.go
package roles

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

func (s *Service) Create(ctx context.Context, req CreateRoleRequest, orgID uuid.UUID) (models.Role, error) {
	role := models.Role{
		OrganizationID: orgID,
		Name:           req.Name,
		Code:           req.Code,
		Description:    req.Description,
		IsActive:       true,
	}
	return s.repo.Create(ctx, role)
}

func (s *Service) List(ctx context.Context, p utils.Pagination, f RoleFilter) ([]models.Role, int64, error) {
	return s.repo.List(ctx, p.Page, p.Size, f)
}

func (s *Service) GetByID(ctx context.Context, id int64, orgID uuid.UUID) (models.Role, error) {
	return s.repo.GetByID(ctx, id, orgID)
}

func (s *Service) Update(ctx context.Context, id int64, req UpdateRoleRequest, orgID uuid.UUID) (models.Role, error) {
	role, err := s.repo.GetByID(ctx, id, orgID)
	if err != nil {
		return role, err
	}

	role.Name = req.Name
	role.Description = req.Description

	return s.repo.Update(ctx, role)
}

func (s *Service) UpdateStatus(ctx context.Context, id int64, isActive bool) error {
	return s.repo.UpdateStatus(ctx, id, isActive)
}

func (s *Service) Delete(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}
