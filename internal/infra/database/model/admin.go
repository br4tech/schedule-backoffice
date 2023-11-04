package model

import (
	"gorm.io/gorm"
)

type UserAdmin struct {
	gorm.Model

	User User `gorm:"polymorphic:User;"`
}
