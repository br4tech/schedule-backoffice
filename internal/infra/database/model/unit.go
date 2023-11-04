package model

import "gorm.io/gorm"

type Unit struct {
	gorm.Model

	UnitID      uint   `gorm:"column:unit_id"`
	Name        string `gorm:"column:name"`
	Location    string `gorm:"column:location"`
	Description string `gorm:"column:description"`
	Rooms       []Room `gorm:"foreignKey:RoomID"`
}
