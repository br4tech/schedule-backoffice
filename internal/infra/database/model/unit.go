package model

import "gorm.io/gorm"

type Unit struct {
	gorm.Model

	Name        string
	Location    string
	Description string
	Rooms       []Room
}
