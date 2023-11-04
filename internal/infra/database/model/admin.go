package model

import (
	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model

	UserID uint `gorm:"column:user_id"`
	User   User `gorm:"foreignKey:UserID"`
}
