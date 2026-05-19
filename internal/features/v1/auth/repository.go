package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/moshfiq123456/ums-backend/internal/models"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) FindUserByID(ctx context.Context, id uuid.UUID) (models.User, error) {
	var user models.User
	err := r.db.WithContext(ctx).
		Preload("Organization").
		Where("id = ? AND deleted_at IS NULL", id).
		First(&user).Error
	return user, err
}

func (r *Repository) FindUserByEmail(ctx context.Context, email string) (models.User, error) {
	var user models.User
	err := r.db.WithContext(ctx).
		Preload("Organization").
		Where("email = ? AND deleted_at IS NULL", email).
		First(&user).Error
	return user, err
}

func (r *Repository) CreateSession(ctx context.Context, session models.LoginSession) error {
	return r.db.WithContext(ctx).Create(&session).Error
}

func (r *Repository) FindSessionByID(ctx context.Context, id uuid.UUID) (models.LoginSession, error) {
	var session models.LoginSession
	err := r.db.WithContext(ctx).
		Where("id = ?", id).
		Where("logged_out_at IS NULL").
		First(&session).Error
	return session, err
}

func (r *Repository) FindOrgByID(ctx context.Context, id uuid.UUID) (models.Organization, error) {
	var org models.Organization
	err := r.db.WithContext(ctx).Where("id = ? AND deleted_at IS NULL", id).First(&org).Error
	return org, err
}

func (r *Repository) FindOrCreateGuestUser(ctx context.Context, orgID uuid.UUID, email *string, name *string) (models.User, bool, error) {
	var user models.User

	if email != nil && *email != "" {
		err := r.db.WithContext(ctx).
			Preload("Organization").
			Where("email = ? AND organization_id = ? AND user_type = 'guest' AND deleted_at IS NULL", *email, orgID).
			First(&user).Error
		if err == nil {
			return user, false, nil
		}
	}

	guestEmail := fmt.Sprintf("guest_%s@ums.guest", uuid.New().String())
	if email != nil && *email != "" {
		guestEmail = *email
	}
	guestName := "Guest"
	if name != nil && *name != "" {
		guestName = *name
	}

	user = models.User{
		OrganizationID: orgID,
		Name:           guestName,
		Email:          guestEmail,
		PasswordHash:   uuid.New().String(),
		UserType:       "guest",
		Status:         "active",
	}
	err := r.db.WithContext(ctx).Create(&user).Error
	if err != nil {
		return user, false, err
	}

	r.db.WithContext(ctx).Preload("Organization").First(&user, "id = ?", user.ID)
	return user, true, nil
}

func (r *Repository) LogoutSession(ctx context.Context, id uuid.UUID) error {
	now := time.Now()
	return r.db.WithContext(ctx).
		Model(&models.LoginSession{}).
		Where("id = ?", id).
		Update("logged_out_at", now).Error
}
