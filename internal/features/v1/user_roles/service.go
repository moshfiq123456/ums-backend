package user_roles

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

func (s *Service) AssignRoles(ctx context.Context, userID uuid.UUID, roleIDs []uint) error {
	return s.repo.AssignRoles(ctx, userID, roleIDs)
}

func (s *Service) RemoveRoles(ctx context.Context, userID uuid.UUID, roleIDs []uint) error {
	return s.repo.RemoveRoles(ctx, userID, roleIDs)
}

func (s *Service) ListRoles(
	ctx context.Context,
	userID uuid.UUID,
	p utils.Pagination,
) ([]models.Role, error) {

	return s.repo.ListRoles(ctx, userID, p.Page, p.Size)
}

