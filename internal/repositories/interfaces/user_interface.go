package repositories

import "github.com/br4tech/schedule-backoffice/internal/domain/entities"

type IUserRepository interface {
	 CreateUser(user *entities.User) error
	 FindByID(id int)(*entities.User, error)
}