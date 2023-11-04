package model

import "gorm.io/gorm"

type Client struct {
	gorm.Model

	User      User       `gorm:"polymorphic:User;"`
	Contracts []Contract `gorm:"foreignKey:ContractID"`
}
