package users

import (
	"context"
	"errors"
	"sync"

	"github.com/google/uuid"
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

// CREATE USER
func (s *Service) Create(ctx context.Context, req CreateUserRequest) (models.User, error) {
	// 1️⃣ Validate fields
	if err := utils.Validate.Struct(req); err != nil {
		return models.User{}, errors.New("validation failed")
	}

	// 2️⃣ Check unique email
	existing, _ := s.repo.GetByEmail(ctx, req.Email)
	if existing.ID != uuid.Nil {
		return models.User{}, errors.New("email already exists")
	}

	// 3️⃣ Hash password
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

// UPDATE USER
func (s *Service) Update(ctx context.Context, id string, req UpdateUserRequest) (models.User, error) {
	// Validate optional fields
	if err := utils.Validate.Struct(req); err != nil {
		return models.User{}, errors.New("validation failed")
	}

	// Update in repo
	return s.repo.Update(ctx, id, req.Name, req.Phone)
}

// LIST
func (s *Service) List(ctx context.Context, p utils.Pagination) ([]models.User, error) {
	return s.repo.List(ctx, p.Page, p.Size)
}

// UserDetail bundles a user with their hierarchy info
type UserDetail struct {
	User     models.User
	Parent   *models.User
	Children []models.User
}

// GET BY ID with all relations (parent & children fetched concurrently)
func (s *Service) GetByID(ctx context.Context, id string) (UserDetail, error) {
	user, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return UserDetail{}, err
	}

	var (
		parent   *models.User
		children []models.User
		wg       sync.WaitGroup
	)

	wg.Add(2)

	go func() {
		defer wg.Done()
		parent, _ = s.repo.GetParent(ctx, user.ID)
	}()

	go func() {
		defer wg.Done()
		children, _ = s.repo.GetChildren(ctx, user.ID)
	}()

	wg.Wait()

	return UserDetail{
		User:     user,
		Parent:   parent,
		Children: children,
	}, nil
}

// DELETE
func (s *Service) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}

// UPDATE STATUS
func (s *Service) UpdateStatus(ctx context.Context, id string, status string) error {
	return s.repo.UpdateStatus(ctx, id, status)
}
