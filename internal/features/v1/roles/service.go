// internal/features/v1/roles/service.go
package roles

import (
	"context"

	"github.com/moshfiq123456/ums-backend/internal/models"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(ctx context.Context, req CreateRoleRequest) (models.Role, error) {
	role := models.Role{
		Name:        req.Name,
		Code:        req.Code,
		Description: req.Description,
		IsActive:    true,
	}
	return s.repo.Create(ctx, role)
}

func (s *Service) List(ctx context.Context) ([]models.Role, error) {
	return s.repo.List(ctx)
}

func (s *Service) GetByID(ctx context.Context, id int64) (models.Role, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *Service) Update(ctx context.Context, id int64, req UpdateRoleRequest) (models.Role, error) {
	role, err := s.repo.GetByID(ctx, id)
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
