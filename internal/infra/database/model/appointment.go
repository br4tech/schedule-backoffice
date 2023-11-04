package model

import (
	"time"

	"github.com/br4tech/schedule-backoffice/internal/infra/database/model"
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

	Weekdays        []Weekday       `gorm:"type:varchar(255)[]"`
	RoomID          uint            `gorm:"column:has_parking"`
	Room            model.Room      `gorm:"foreignKey:RoomID"`
	HourlyRate      float64         `gorm:"column:hourly_rate"`
	Period          string          `gorm:"column:period"`
	StartAt         time.Time       `gorm:"column:end_at"`
	EndAt           time.Time       `gorm:"column:start_at"`
	AppointmentType AppointmentType `gorm:"column:appointment_type"`
	ContractID      uint            `gorm:"column:contract_id"`
	Contract        Contract        `gorm:"foreignKey:ContractID"`
	HasParking      bool            `gorm:"column:has_parking"`
}
