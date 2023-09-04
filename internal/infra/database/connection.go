package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitializeDatabase() (*gorm.DB, error) {
	  dsn := "host=localhost user=guilherme password=admin dbname=pronto_development port=5432 sslmode=disable" // Substitua com suas configurações

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

    if err != nil {
        return nil, err
    }

    DB = db

    return db, nil
}

func CloseDatabase() {
    sqlDB, _ := DB.DB()
    sqlDB.Close()
}