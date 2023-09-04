package application

import "github.com/br4tech/schedule-backoffice/internal/domain/entities"

type IUserApplication interface {
	FindByID(id int)(*entities.User, error)
}