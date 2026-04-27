package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Organization struct {
	ID        uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name      string         `gorm:"size:200;not null"`
	Slug      string         `gorm:"size:100;uniqueIndex;not null"`
	Domain    *string        `gorm:"size:200"`
	LogoURL   *string        `gorm:"size:500"`
	IsActive  bool           `gorm:"default:true"`
	Plan      string         `gorm:"size:50;default:free"`
	Settings  string         `gorm:"type:jsonb;default:'{}'"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Users          []User          `gorm:"foreignKey:OrganizationID"`
	Roles          []Role          `gorm:"foreignKey:OrganizationID"`
	ServiceClients []ServiceClient `gorm:"foreignKey:OrganizationID"`
}
