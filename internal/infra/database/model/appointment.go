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

	Weekdays        []Weekday       `gorm:"type:varchar(255)"`
	RoomID          int             `gorm:"column:has_parking"`
	Room            Room            `gorm:"foreignKey:RoomID"`
	HourlyRate      float64         `gorm:"column:hourly_rate"`
	Period          string          `gorm:"column:period"`
	StartTimeAt     time.Duration   `gorm:"column:end_time_at"`
	EndTimeAt       time.Duration   `gorm:"column:start_time_at"`
	AppointmentType AppointmentType `gorm:"column:appointment_type"`
	ContractID      int             `gorm:"column:contract_id"`
	Contract        Contract        `gorm:"foreignKey:ContractID"`
	HasParking      bool            `gorm:"column:has_parking"`
}
