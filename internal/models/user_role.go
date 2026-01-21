package models

import "github.com/google/uuid"

type UserRole struct {
	ID     uint      `gorm:"primaryKey"`
	UserID uuid.UUID `gorm:"type:uuid;not null;index"`
	RoleID uint      `gorm:"not null;index"`

	User User `gorm:"foreignKey:UserID"`
	Role Role `gorm:"foreignKey:RoleID"`
}
