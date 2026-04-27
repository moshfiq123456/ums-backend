package models

import "time"

type Permission struct {
	ID          uint      `gorm:"primaryKey"`
	Code        string    `gorm:"size:100;uniqueIndex;not null"`
	Name        string    `gorm:"size:100;not null"`
	Description string    `gorm:"size:255"`
	Service     string    `gorm:"size:50;default:ums;index"` // ums | ecommerce | admin
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}

