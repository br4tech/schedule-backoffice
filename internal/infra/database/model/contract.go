package model

import (
	"time"

	"gorm.io/gorm"
)

type Contract struct {
	gorm.Model

	StartAt        time.Time     `gorm:"column:start_at"`
	EndAt          time.Time     `gorm:"column:end_at"`
	Amount         float64       `gorm:"column:amount"`
	DueAt          time.Time     `gorm:"column:due_at"`
	RevenueAt      time.Time     `gorm:"column:revenue_at"`
	Forfeit        string        `gorm:"column:forfeit"`
	Active         bool          `gorm:"column:active"`
	ClientID       uint          `gorm:"column:client_id"`
	Client         Client        `gorm:"foreignKey:ClientID"`
	ContractTypeID uint          `gorm:"foreignKey:ContractTypeID"`
	ContracType    ContractType  `gorm:"foreignKey:ContractTypeID"`
	CompanyID      uint          `gorm:"column:company_id"`
	Company        Company       `gorm:"foreignKey:CompanyID"`
	Appointments   []Appointment `gorm:"foreignKey:AppointmentID"`
}
