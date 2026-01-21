package models

import (
	"time"

	"github.com/google/uuid"
)

type UserHierarchy struct {
	ID           uint      `gorm:"primaryKey"`
	ParentUserID uuid.UUID `gorm:"type:uuid;not null;index"`
	ChildUserID  uuid.UUID `gorm:"type:uuid;not null;index"`
	CreatedAt    time.Time

	ParentUser User `gorm:"foreignKey:ParentUserID"`
	ChildUser  User `gorm:"foreignKey:ChildUserID"`
}

func (UserHierarchy) TableName() string {
	return "user_hierarchy"
}