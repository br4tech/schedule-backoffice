package model

import (
	"gorm.io/gorm"
)

type Secretary struct {
	gorm.Model

	User          User        `gorm:"polymorphic:User;"`
	DepartamentID uint        `gorm:"column:departament_id"`
	Departament   Departament `gorm:"foreignKey:DepartamentID"`
}
