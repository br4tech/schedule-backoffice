package main

import (
	"fmt"

	"github.com/br4tech/schedule-backoffice/internal/infra/database"
	"github.com/br4tech/schedule-backoffice/internal/infra/repository"
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
	userRepository := repository.NewUserRepository(adapter)
	user, err := userRepository.FindAll()
	if err != nil {
		panic("Ocorreu um erro ao usar o user repository")
	}

	fmt.Println(user)
}
