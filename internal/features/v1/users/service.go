package users

import (
	"context"

	"github.com/moshfiq123456/ums-backend/internal/models"
	"github.com/moshfiq123456/ums-backend/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repo *UserRepository
}

func NewService(repo *UserRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(ctx context.Context, req CreateUserRequest) (models.User, error) {
	hash, _ := bcrypt.GenerateFromPassword([]byte(req.Password), 12)

	user := models.User{
		Name:         req.Name,
		Email:        req.Email,
		PasswordHash: string(hash),
		Phone:        req.Phone,
		Status:       "active",
	}

	return s.repo.Create(ctx, user)
}

func (s *Service) List(ctx context.Context, p utils.Pagination) ([]models.User, error) {
	return s.repo.List(ctx, p.Page, p.Size)
}

func (s *Service) GetByID(ctx context.Context, id string) (models.User, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *Service) Update(ctx context.Context, id string, req UpdateUserRequest) (models.User, error) {
	return s.repo.Update(ctx, id, req.Name, req.Phone)
}

func (s *Service) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}

func (s *Service) UpdateStatus(ctx context.Context, id string, status string) error {
	return s.repo.UpdateStatus(ctx, id, status)
}
