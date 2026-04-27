package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type ServiceClient struct {
	ID                uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	OrganizationID    uuid.UUID      `gorm:"type:uuid;not null;index"`
	Name              string         `gorm:"size:100;not null"`
	ClientID          string         `gorm:"size:100;uniqueIndex;not null"`
	ClientSecretHash  string         `gorm:"not null"`
	AllowedScopes     pq.StringArray `gorm:"type:text[]"`
	IsActive          bool           `gorm:"default:true"`
	CreatedAt         time.Time
	UpdatedAt         time.Time

	Organization Organization `gorm:"foreignKey:OrganizationID"`
}
