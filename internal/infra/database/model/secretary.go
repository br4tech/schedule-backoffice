package model

import (
	"github.com/br4tech/schedule-backoffice/internal/infra/database/model"
	"gorm.io/gorm"
)

type Secretary struct {
	gorm.Model

	UserID      uint       `gorm:"column:user_id"`
	User        model.User `gorm:"foreignKey:UserID"`
	Departament string     `gorm:"column:departament"`
}
