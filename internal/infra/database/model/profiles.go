package model

import "gorm.io/gorm"

type Profile struct {
	gorm.Model

	Name string `gorm:"column:name"`
}
