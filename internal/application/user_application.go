package application

import (
	application "github.com/br4tech/schedule-backoffice/internal/application/interfaces"
	"github.com/br4tech/schedule-backoffice/internal/domain/entities"
	services "github.com/br4tech/schedule-backoffice/internal/domain/services/interface"
)

type UserApplication struct {
	userService services.IUserService
}

func NewUserApplication(userService services.IUserService) application.IUserApplication {
	return &UserApplication{
		userService: userService,
	}
}

func(app *UserApplication) FindByID(id int)(*entities.User, error){
	return app.userService.FindByID(id)
}