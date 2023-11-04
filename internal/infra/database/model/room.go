package model

import (
	"gorm.io/gorm"
)

type Room struct {
	gorm.Model

	RoomID uint    `gorm:"column:room_id"`
	Number int     `gorm:"column:number"`
	Size   float64 `gorm:"column:size"`
	Price  float64 `gorm:"column:price"`
	Color  string  `gorm:"column:color"`
	UnitID uint    `gorm:"column:unit_id"`
	Unit   Unit    `gorm:"foreignKey:unitID"`
}
