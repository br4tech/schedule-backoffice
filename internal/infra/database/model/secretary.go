package model

import (
	"gorm.io/gorm"
)

type Secretary struct {
	gorm.Model

	User        User   `gorm:"polymorphic:User;"`
	Departament string `gorm:"column:departament"`
}
