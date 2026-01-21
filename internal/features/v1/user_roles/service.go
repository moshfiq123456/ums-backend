package user_roles

import (
	"context"

	"github.com/google/uuid"
	"github.com/moshfiq123456/ums-backend/internal/models"
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

func (s *Service) ListRoles(ctx context.Context, userID uuid.UUID) ([]models.Role, error) {
	// Use models.Role instead of undefined Role
	return s.repo.ListRoles(ctx, userID)
}
