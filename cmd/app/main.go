package main

import (
	"github.com/br4tech/schedule-backoffice/internal/infra/database"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := ""
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Falha ao conectar ao banco de dados")
	}

	adapter := database.NewGormAdapter(db)
	db = adapter.GetDB()
}
