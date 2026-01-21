package permissions

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

func (s *Service) Create(ctx context.Context, req CreatePermissionRequest) error {
	p := models.Permission{
		Code:        req.Code,
		Name:        req.Name,
		Description: req.Description,
	}
	return s.repo.Create(ctx, p)
}

func (s *Service) List(ctx context.Context) ([]models.Permission, error) {
	return s.repo.List(ctx)
}

func (s *Service) Get(ctx context.Context, id uint) (models.Permission, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *Service) Update(ctx context.Context, id uint, req UpdatePermissionRequest) (models.Permission, error) {
	p, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return p, err
	}

	p.Name = req.Name
	p.Description = req.Description

	return p, s.repo.Update(ctx, p)
}

func (s *Service) Delete(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}
