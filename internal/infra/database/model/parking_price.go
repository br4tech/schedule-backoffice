package model

import "gorm.io/gorm"

type ParkingPrice struct {
	gorm.Model

	AppointmentID uint    `gorm:"column:appointment_id"`
	Price         float64 `gorm:"column:price"`
}
