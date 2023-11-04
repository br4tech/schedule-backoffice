package model

import (
	"github.com/br4tech/schedule-backoffice/internal/infra/database/model"
	"gorm.io/gorm"
)

type Room struct {
	gorm.Model

	Number int        `gorm:"column:number"`
	Size   float64    `gorm:"column:size"`
	Price  float64    `gorm:"column:price"`
	Color  string     `gorm:"column:color"`
	UnitID uint       `gorm:"column:unit_id"`
	Unit   model.Unit `gorm:"foreignKey:unitID"`
}
