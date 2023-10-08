package model

import "gorm.io/gorm"

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

	Weekdays   []Weekday `gorm:"type:varchar(255)[]"`
	RoomID     uint
	HourlyRate float64
	ContractID uint
	Contract   Contract `gorm:"foreignKey:ContractID"`
	HasParking bool     `gorm:"column:has_parking"`
}
