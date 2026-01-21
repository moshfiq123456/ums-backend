package models

import "github.com/google/uuid"

type UserPermission struct {
	ID           uint      `gorm:"primaryKey"`
	UserID       uuid.UUID `gorm:"type:uuid;not null;index"`
	PermissionID uint      `gorm:"not null;index"`
	Allow        bool      `gorm:"default:true"`

	User       User       `gorm:"foreignKey:UserID"`
	Permission Permission `gorm:"foreignKey:PermissionID"`
}
