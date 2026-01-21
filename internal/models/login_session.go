package models

import (
	"time"

	"github.com/google/uuid"
)

// internal/models/login_session.go
type LoginSession struct {
	ID                 uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID             uuid.UUID  `gorm:"type:uuid;not null;index"`

	RefreshTokenHash   string     `gorm:"size:255;not null;index"`
	RefreshExpiresAt   time.Time  `gorm:"not null"`

	IPAddress          *string    `gorm:"size:45"`
	UserAgent          *string    `gorm:"type:text"`

	LoggedInAt         time.Time  `gorm:"autoCreateTime"`
	LoggedOutAt        *time.Time

	User User `gorm:"foreignKey:UserID"`
}
