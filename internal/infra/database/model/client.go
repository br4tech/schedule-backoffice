package model

import "gorm.io/gorm"

type Client struct {
	gorm.Model

	User
	Contracts []ClientContract
}
