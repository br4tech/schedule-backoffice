package model

import "gorm.io/gorm"

type ParkingPrice struct {
	gorm.Model

	Price         float64     `gorm:"column:price"`
	AppointmentID int         `gorm:"column:appointment_id"`
	Appointment   Appointment `gorm:"foreignKey:AppointmentID"`
}
