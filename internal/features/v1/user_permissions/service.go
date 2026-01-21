package user_permissions

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

func (s *Service) AssignPermissions(ctx context.Context, userID uuid.UUID, permissionIDs []uint) error {
	return s.repo.AssignPermissions(ctx, userID, permissionIDs)
}

func (s *Service) RemovePermissions(ctx context.Context, userID uuid.UUID, permissionIDs []uint) error {
	return s.repo.RemovePermissions(ctx, userID, permissionIDs)
}

func (s *Service) ListPermissions(ctx context.Context, userID uuid.UUID) ([]models.Permission, error) {
	return s.repo.ListPermissions(ctx, userID)
}
