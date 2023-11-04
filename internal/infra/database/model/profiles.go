package model

import "gorm.io/gorm"

type Profile struct {
	gorm.Model

	ProfileID uint   `gorm:"column:profile_id"`
	Name      string `gorm:"column:name"`
}
