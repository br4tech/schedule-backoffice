package model

import "gorm.io/gorm"

type ContractType struct {
	gorm.Model

	ContractTypeID uint       `gorm:"column:contract_type_id"`
	Name           string     `gorm:"column:name"`
	Contract       []Contract `gorm:"foreignKey:ContractID"`
}
