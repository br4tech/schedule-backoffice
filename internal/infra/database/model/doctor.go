package model

import "gorm.io/gorm"

type Doctor struct {
	gorm.Model

	Use            User    `gorm:"foreignKey:UserID"`
	Crm            string  `gorm:"column:crm"`
	Specialty      string  `gorm:"column:speciality"`
	ReturnDuration int     `gorm:"column:return_duration"`
	AcceptsReturns bool    `gorm:"column:accept_return"`
	Notes          string  `gorm:"column:notes"`
	Supervisor     *Doctor `gorm:"foreignKey:SupervisorID"`
	SupervisorID   uint
}
