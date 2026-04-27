package models

import (
	"time"

	"github.com/google/uuid"
)

type Role struct {
	ID             int64     `gorm:"primaryKey;autoIncrement"`
	OrganizationID uuid.UUID `gorm:"type:uuid;not null;index"`
	Name           string    `gorm:"size:50;not null"`
	Code           string    `gorm:"size:100;not null"`
	Description    string    `gorm:"type:text"`
	IsActive       bool      `gorm:"default:true"`
	CreatedAt      time.Time
	UpdatedAt      time.Time

	Organization Organization `gorm:"foreignKey:OrganizationID"`
}
