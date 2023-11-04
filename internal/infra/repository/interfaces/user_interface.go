package repository

import "github.com/br4tech/schedule-backoffice/internal/domain/entity"

type IUserRepository interface {
	 CreateUser(user *entity.User) error
	 FindByID(id int)(*entity.User, error)
}