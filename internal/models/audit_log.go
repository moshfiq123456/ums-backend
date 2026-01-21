package models

import (
	"time"

	"github.com/google/uuid"
)

type AuditLog struct {
	ID        int64     `gorm:"primaryKey;autoIncrement"`
	UserID    *uuid.UUID `gorm:"type:uuid;index"`
	Action    string    `gorm:"size:100;not null"`
	Entity    string    `gorm:"size:50;not null"`
	EntityID  *string   `gorm:"size:50"`
	OldValue  []byte    `gorm:"type:jsonb"`
	NewValue  []byte    `gorm:"type:jsonb"`
	CreatedAt time.Time `gorm:"autoCreateTime"`

	User *User `gorm:"foreignKey:UserID"`
}