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
	SupervisorID   uint    `gorm:"column:supervisor_id"`
	Supervisor     *Doctor `gorm:"foreignKey:SupervisorID"`
}
