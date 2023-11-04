package main

import (
	"github.com/br4tech/schedule-backoffice/internal/application"
	"github.com/br4tech/schedule-backoffice/internal/domain/service"
	"github.com/br4tech/schedule-backoffice/internal/infra/database"
	"github.com/br4tech/schedule-backoffice/internal/infra/repository"
)

func main(){
  _, err := database.InitializeDatabase()
	if err != nil {
		return
	}

	userRepo := repository.NewUserRepository()
	userService := service.NewUserService(userRepo)
	userApplication := application.NewApplication(userService)

	defer database.CloseDatabase()
}