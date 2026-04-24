package users

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/moshfiq123456/ums-backend/internal/models"
	"github.com/moshfiq123456/ums-backend/internal/utils"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/sync/errgroup"
)

type Service struct {
	repo *UserRepository
}

func NewService(repo *UserRepository) *Service {
	return &Service{repo: repo}
}

// CREATE USER
func (s *Service) Create(ctx context.Context, req CreateUserRequest) (models.User, error) {
	if err := utils.Validate.Struct(req); err != nil {
		return models.User{}, errors.New("validation failed")
	}

	// Email check and bcrypt are independent — run them concurrently.
	var hash []byte

	g, gctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		existing, _ := s.repo.GetByEmail(gctx, req.Email)
		if existing.ID != uuid.Nil {
			return errors.New("email already exists")
		}
		return nil
	})

	g.Go(func() error {
		var err error
		hash, err = bcrypt.GenerateFromPassword([]byte(req.Password), 12)
		return err
	})

	if err := g.Wait(); err != nil {
		return models.User{}, err
	}

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
	)

	g, gctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		var err error
		parent, err = s.repo.GetParent(gctx, user.ID)
		return err
	})

	g.Go(func() error {
		var err error
		children, err = s.repo.GetChildren(gctx, user.ID)
		return err
	})

	if err := g.Wait(); err != nil {
		return UserDetail{}, err
	}

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
