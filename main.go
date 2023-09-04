package main

import (
	"github.com/br4tech/schedule-backoffice/internal/infra/database"
	"github.com/br4tech/schedule-backoffice/internal/infra/repositories"
)

func main(){
  _, err := database.InitializeDatabase()
	if err != nil {
		return
	}

	userRepo := repositories.NewUserRepository()

	defer database.CloseDatabase()
}