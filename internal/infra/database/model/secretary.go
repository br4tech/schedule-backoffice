package model

import (
	"gorm.io/gorm"
)

type Secretary struct {
	gorm.Model

	User          User        `gorm:"polymorphic:Role;"`
	DepartamentID uint        `gorm:"column:departament_id"`
	Departament   Departament `gorm:"foreignKey:DepartamentID"`
}
