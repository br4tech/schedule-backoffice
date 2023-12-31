package model

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Name         string       `gorm:"column:name"`
	Email        string       `gorm:"column:email"`
	Password     string       `gorm:"column:password"`
	Phone        string       `gorm:"column:phone"`
	IsIndividual bool         `gorm:"column:is_individual"`
	Document     string       `gorm:"column:document"`
	Active       bool         `gorm:"column:active"`
	Overdue      bool         `gorm:"column:overdue"`
	Role         string       `gorm:"column:role"`
	RoleID       int          `gorm:"column:role_id"`
	RoleType     string       `gorm:"column:role_type"`
	ProfileID    uint         `gorm:"column:profile_id"`
	Profiles     Profile      `gorm:"foreignKey:ProfileID"`
	Permission   []Permission `gorm:"many2many:user_permissions;"`
}
