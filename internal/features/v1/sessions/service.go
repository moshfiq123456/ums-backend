package sessions

import (
	"context"

	"github.com/google/uuid"
	"github.com/moshfiq123456/ums-backend/internal/utils"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) ListActive(ctx context.Context, p utils.Pagination, f ActiveSessionFilter) ([]ActiveSessionResponse, int64, error) {
	sessions, total, err := s.repo.ListActive(ctx, p.Page, p.Size, f)
	if err != nil {
		return nil, 0, err
	}

	resp := make([]ActiveSessionResponse, 0, len(sessions))
	for _, s := range sessions {
		resp = append(resp, toResponse(s))
	}
	return resp, total, nil
}

func (s *Service) ForceLogout(ctx context.Context, sessionID uuid.UUID, orgID uuid.UUID) error {
	return s.repo.ForceLogout(ctx, sessionID, orgID)
}
