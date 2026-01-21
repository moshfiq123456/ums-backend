package models

type RolePermission struct {
	ID           uint `gorm:"primaryKey"`
	RoleID       uint `gorm:"not null;index"`
	PermissionID uint `gorm:"not null;index"`

	Role       Role       `gorm:"foreignKey:RoleID"`
	Permission Permission `gorm:"foreignKey:PermissionID"`
}
