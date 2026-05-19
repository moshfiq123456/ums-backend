package models

import (
	"time"

	"github.com/google/uuid"
)

type Permission struct {
	ID             uint      `gorm:"primaryKey"`
	OrganizationID uuid.UUID `gorm:"type:uuid;not null;index"`
	Code           string    `gorm:"size:100;not null"`
	Name           string    `gorm:"size:100;not null"`
	Description    string    `gorm:"size:255"`
	CreatedAt      time.Time `gorm:"autoCreateTime"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime"`

	Organization Organization `gorm:"foreignKey:OrganizationID"`
}

