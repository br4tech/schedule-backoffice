package model

import "gorm.io/gorm"

type ParkingPrice struct {
	gorm.Model

	ParkingPriceID uint    `gorm:"column:parking_price_id"`
	AppointmentID  uint    `gorm:"column:appointment_id"`
	Price          float64 `gorm:"column:price"`
}
