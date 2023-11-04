package model

import (
	"time"

	"gorm.io/gorm"
)

type AppointmentType string

const (
	Diario    AppointmentType = "Diario"
	Quinzenal AppointmentType = "Quinzenal"
)

type Weekday string

const (
	Domingo Weekday = "Domingo"
	Segunda Weekday = "Segunda"
	Terca   Weekday = "Terca"
	Quarta  Weekday = "Quarta"
	Quinta  Weekday = "Quinta"
	Sexta   Weekday = "Sexta"
	Sabado  Weekday = "Sabado"
)

type Appointment struct {
	gorm.Model

	AppointmentID   uint            `gorm:"foreignKey:AppointmentID"`
	Weekdays        []Weekday       `gorm:"type:varchar(255)[]"`
	RoomID          uint            `gorm:"column:has_parking"`
	Room            Room            `gorm:"foreignKey:RoomID"`
	HourlyRate      float64         `gorm:"column:hourly_rate"`
	Period          string          `gorm:"column:period"`
	StartTimeAt     time.Ticker     `gorm:"column:end_at"`
	EndTimeAt       time.Ticker     `gorm:"column:start_at"`
	AppointmentType AppointmentType `gorm:"column:appointment_type"`
	ContractID      uint            `gorm:"column:contract_id"`
	Contract        Contract        `gorm:"foreignKey:ContractID"`
	HasParking      bool            `gorm:"column:has_parking"`
}
