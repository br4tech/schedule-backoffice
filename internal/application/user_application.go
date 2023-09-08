package application

import (
	"github.com/br4tech/schedule-backoffice/internal/domain/entity"
	service "github.com/br4tech/schedule-backoffice/internal/domain/service/interface"
)

type Application struct {
    userService service.IUserService
}

func NewApplication(userService service.IUserService) *Application {
    return &Application{
        userService: userService,
    }
}

func (app *Application) GetUserByID(id int) (*entity.User, error) {
    return app.userService.GetUserByID(id)
}