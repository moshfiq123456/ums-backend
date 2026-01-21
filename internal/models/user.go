package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID              uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name            string         `gorm:"size:100;not null"`
	Email           string         `gorm:"size:150;uniqueIndex;not null"`
	PasswordHash    string         `gorm:"not null"`
	Phone           *string        `gorm:"size:20"`
	Status          string         `gorm:"size:20;default:active"`
	EmailVerifiedAt *time.Time
	LastLoginAt     *time.Time

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	// Relations
	Roles       []Role           `gorm:"many2many:user_roles"`
	Permissions []UserPermission `gorm:"foreignKey:UserID"`
}
