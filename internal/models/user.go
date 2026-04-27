package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID             uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	OrganizationID uuid.UUID      `gorm:"type:uuid;not null;index"`
	Name           string         `gorm:"size:100;not null"`
	Email          string         `gorm:"size:150;not null"`
	PasswordHash   string         `gorm:"not null"`
	Phone          *string        `gorm:"size:20"`
	AvatarURL      *string        `gorm:"size:500"`
	UserType       string         `gorm:"size:20;default:member"` // admin | customer | member
	Status         string         `gorm:"size:20;default:active"`
	Metadata       string         `gorm:"type:jsonb;default:'{}'"`
	EmailVerifiedAt *time.Time
	LastLoginAt     *time.Time

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	// Relations
	Organization Organization     `gorm:"foreignKey:OrganizationID"`
	Roles        []Role           `gorm:"many2many:user_roles"`
	Permissions  []UserPermission `gorm:"foreignKey:UserID"`
}
