package user_permissions

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

func (s *Service) AssignPermissions(ctx context.Context, userID uuid.UUID, orgID uuid.UUID, permissionIDs []uint) error {
	return s.repo.AssignPermissions(ctx, userID, orgID, permissionIDs)
}

func (s *Service) RemovePermissions(ctx context.Context, userID uuid.UUID, orgID uuid.UUID, permissionIDs []uint) error {
	return s.repo.RemovePermissions(ctx, userID, orgID, permissionIDs)
}

func (s *Service) ListPermissions(ctx context.Context, userID uuid.UUID, orgID uuid.UUID, p utils.Pagination) ([]models.Permission, error) {
	return s.repo.ListPermissions(ctx, userID, orgID, p.Page, p.Size)
}

func (s *Service) ListAll(ctx context.Context, p utils.Pagination, f UserPermissionFilter) ([]UserPermissionDetail, int64, error) {
	return s.repo.ListAll(ctx, p.Page, p.Size, f)
}

