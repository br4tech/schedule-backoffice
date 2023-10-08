package model

import "gorm.io/gorm"

type Room struct {
	gorm.Model

	Number int     `gorm:"column:number"`
	Size   float64 `gorm:"column:size"`
	Price  float64 `gorm:"column:price"`
	Color  string  `gorm:"column:color"`
	UnitID uint
}
