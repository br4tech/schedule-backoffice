package model

import "gorm.io/gorm"

type Unit struct {
	gorm.Model

	Name        string `gorm:"column:name"`
	Location    string `gorm:"column:location"`
	Description string `gorm:"column:description"`
	Rooms       []Room `gorm:"many2many:rooms;"`
}
