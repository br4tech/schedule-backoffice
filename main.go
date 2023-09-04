package main

import (
	"github.com/br4tech/schedule-backoffice/internal/application"
	"github.com/br4tech/schedule-backoffice/internal/domain/services"
	"github.com/br4tech/schedule-backoffice/internal/infra/database"
	"github.com/br4tech/schedule-backoffice/internal/infra/repositories"
)

func main(){
  _, err := database.InitializeDatabase()
	if err != nil {
		return
	}

	userRepo := repositories.NewUserRepository()
	userService := services.NewUserService(userRepo)
	userApplication := application.NewUserApplication(userService)

	defer database.CloseDatabase()
}