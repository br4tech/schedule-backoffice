package model

import "gorm.io/gorm"

type ClientContract struct {
	gorm.Model

	ClientID   uint `gorm:"column:client_id"`
	ContractID uint `gorm:"column:contract_id"`
}
