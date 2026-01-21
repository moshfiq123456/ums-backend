package user_hierarchy

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) AssignChild(ctx context.Context, parentID, childID uuid.UUID) error {
	if parentID == childID {
		return errors.New("parent and child cannot be same")
	}

	if s.repo.Exists(ctx, parentID, childID) {
		return errors.New("hierarchy already exists")
	}

	return s.repo.Create(ctx, parentID, childID)
}

func (s *Service) RemoveChild(ctx context.Context, parentID, childID uuid.UUID) error {
	return s.repo.Delete(ctx, parentID, childID)
}

func (s *Service) GetChildren(ctx context.Context, userID uuid.UUID) (interface{}, error) {
	users, err := s.repo.GetChildren(ctx, userID)
	if err != nil {
		return nil, err
	}

	resp := make([]UserResponse, 0, len(users))
	for _, u := range users {
		resp = append(resp, toUserResponse(u))
	}

	return resp, nil
}

func (s *Service) GetParent(ctx context.Context, userID uuid.UUID) (interface{}, error) {
	user, err := s.repo.GetParent(ctx, userID)
	if err != nil {
		return nil, err
	}

	return toUserResponse(user), nil
}

func (s *Service) CheckHierarchy(ctx context.Context, parentID, childID uuid.UUID) (interface{}, error) {
	return CheckHierarchyResponse{
		IsRelated: s.repo.Exists(ctx, parentID, childID),
	}, nil
}
