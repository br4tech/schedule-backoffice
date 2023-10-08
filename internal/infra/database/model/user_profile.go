package model

import "gorm.io/gorm"

type UserProfile struct {
	gorm.Model

	ProfileID uint `gorm:"column:profile_id"`
	UserID    uint `gorm:"column:user_id"`
}
