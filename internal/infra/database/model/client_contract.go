package model

import "gorm.io/gorm"

type ClientContract struct {
	gorm.Model

	ClientID   uint
	ContractID uint
}
