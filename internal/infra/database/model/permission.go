package model

import "gorm.io/gorm"

type Permission struct {
	gorm.Model

	PermissionID uint   `gorm:"column:permission_id"`
	Name         string `gorm:"column:name"`
}
