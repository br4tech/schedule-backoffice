package model

import (
	"gorm.io/gorm"
)

type Departament struct {
	gorm.Model

	Name string `gorm:"column:name"`
}
