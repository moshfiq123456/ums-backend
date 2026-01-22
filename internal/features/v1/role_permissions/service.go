package role_permissions

import (
	"context"
	"errors"

	"github.com/moshfiq123456/ums-backend/internal/utils"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Assign(ctx context.Context, roleID, permissionID uint) error {
	if s.repo.Exists(ctx, roleID, permissionID) {
		return errors.New("permission already assigned to role")
	}
	return s.repo.Create(ctx, roleID, permissionID)
}

func (s *Service) BulkAssign(ctx context.Context, roleID uint, permissionIDs []uint) error {
	for _, pid := range permissionIDs {
		if !s.repo.Exists(ctx, roleID, pid) {
			if err := s.repo.Create(ctx, roleID, pid); err != nil {
				return err
			}
		}
	}
	return nil
}

func (s *Service) Remove(ctx context.Context, roleID, permissionID uint) error {
	return s.repo.Delete(ctx, roleID, permissionID)
}

func (s *Service) List(
	ctx context.Context,
	roleID uint,
	p utils.Pagination,
) ([]PermissionResponse, error) {

	perms, err := s.repo.List(ctx, roleID, p.Page, p.Size)
	if err != nil {
		return nil, err
	}

	resp := make([]PermissionResponse, 0, len(perms))
	for _, perm := range perms {
		resp = append(resp, toPermissionResponse(perm))
	}

	return resp, nil
}

