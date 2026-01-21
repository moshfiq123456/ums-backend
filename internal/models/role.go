package models

import "time"

type Role struct {
	ID          int64     `gorm:"primaryKey;autoIncrement"`
	Name        string    `gorm:"size:50;uniqueIndex;not null"`
	Code        string    `gorm:"size:100;uniqueIndex;not null"`
	Description string    `gorm:"type:text"`
	IsActive    bool      `gorm:"default:true"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
