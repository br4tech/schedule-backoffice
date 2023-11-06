package model

import "gorm.io/gorm"

type ContractType struct {
	gorm.Model

	Name string `gorm:"column:name"`
}
