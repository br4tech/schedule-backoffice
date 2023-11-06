package model

import "gorm.io/gorm"

type Company struct {
	gorm.Model

	Document       string `gorm:"column:document"`
	Wallet         string `gorm:"column:wallet"`
	Agency         int    `gorm:"column:agency"`
	CurrentAccount int    `gorm:"column:current_account"`
	Digit          int    `gorm:"column:digit"`
	CompanyCode    string `gorm:"column:company_code"`
	ZipCode        string `gorm:"column:zip_code"`
	City           string `gorm:"column:city"`
	Adress         string `gorm:"column:adress"`
}
