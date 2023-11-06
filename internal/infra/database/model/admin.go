package model

import (
	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model

	User User `gorm:"polymorphic:Role;"`
}
