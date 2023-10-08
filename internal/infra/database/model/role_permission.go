package model

import "gorm.io/gorm"

type RolePermission struct {
	gorm.Model

	PermissionID uint `gorm:"column:permission_id"`
	UserID       uint `gorm:"column:user_id"`
}
