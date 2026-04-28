package organizations

import (
	"context"
	"errors"
	"strings"

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

func (s *Service) Create(ctx context.Context, req CreateOrgRequest) (models.Organization, error) {
	slug := strings.ToLower(strings.ReplaceAll(req.Slug, " ", "-"))
	if s.repo.SlugExists(ctx, slug, nil) {
		return models.Organization{}, errors.New("slug already in use")
	}

	plan := req.Plan
	if plan == "" {
		plan = "free"
	}

	org := models.Organization{
		Name: req.Name,
		Slug: slug,
		Plan: plan,
	}
	if req.Domain != "" {
		org.Domain = &req.Domain
	}

	return s.repo.Create(ctx, org)
}

func (s *Service) List(ctx context.Context, p utils.Pagination, f OrgFilter) ([]models.Organization, int64, error) {
	return s.repo.List(ctx, p.Page, p.Size, f)
}

func (s *Service) GetByID(ctx context.Context, id uuid.UUID) (models.Organization, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *Service) Update(ctx context.Context, id uuid.UUID, req UpdateOrgRequest) (models.Organization, error) {
	org, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return org, err
	}

	if req.Name != "" {
		org.Name = req.Name
	}
	if req.Plan != "" {
		org.Plan = req.Plan
	}
	if req.Domain != "" {
		org.Domain = &req.Domain
	}

	return s.repo.Update(ctx, org)
}

func (s *Service) UpdateStatus(ctx context.Context, id uuid.UUID, isActive bool) error {
	return s.repo.UpdateStatus(ctx, id, isActive)
}

func (s *Service) Delete(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}

func (s *Service) Stats(ctx context.Context, id uuid.UUID) (int64, int64) {
	return s.repo.CountUsers(ctx, id), s.repo.CountRoles(ctx, id)
}
