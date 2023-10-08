package model

import "gorm.io/gorm"

type Secretary struct {
	gorm.Model

	UserID      uint   `gorm:"column:user_id"`
	Departament string `gorm:"column:departament"`
}
