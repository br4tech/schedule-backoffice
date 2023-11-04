package model

import "gorm.io/gorm"

type Doctor struct {
	gorm.Model

	Use            User    `gorm:"polymorphic:User;"`
	Crm            string  `gorm:"column:crm"`
	Specialty      string  `gorm:"column:speciality"`
	AcceptsReturns bool    `gorm:"column:accept_return"`
	ReturnDuration int     `gorm:"column:return_duration"`
	Notes          string  `gorm:"column:notes"`
	SupervisorID   uint    `gorm:"column:supervisor_id"`
	Supervisor     *Doctor `gorm:"foreignKey:SupervisorID"`
}
