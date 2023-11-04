package model

import (
	"gorm.io/gorm"
)

type Departament struct {
	gorm.Model

	DepartamentID uint   `gorm:"column:departament_id"`
	Name          string `gorm:"column:name"`
}
