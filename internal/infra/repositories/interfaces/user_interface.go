package repositories

import "github.com/br4tech/schedule-backoffice/internal/domain/entities"

type IUserRepository interface {
	Create(user *entities.User) error
	FindByID(id int)(*entities.User, error)
}