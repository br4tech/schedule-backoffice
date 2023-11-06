package main

import (
	"github.com/br4tech/schedule-backoffice/internal/infra/database"
	"github.com/br4tech/schedule-backoffice/internal/infra/database/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:admin@tcp(localhost:3306)/schedule_app?charset=utf8mb4&parseTime=True"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Falha ao conectar ao banco de dados")
	}

	adapter := database.NewGormAdapter(db)
	db = adapter.GetDB()

	db.AutoMigrate(
		&model.Company{},
		&model.Profile{},
		&model.Permission{},
		&model.Departament{},
		&model.User{},
		&model.Admin{},
		&model.Doctor{},
		&model.Secretary{},
		&model.Client{},
		&model.ContractType{},
		&model.Contract{},
		&model.Unit{},
		&model.Room{},
		&model.Appointment{},
		&model.ParkingPrice{},
	)
}
