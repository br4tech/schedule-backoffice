package repository

import "github.com/br4tech/schedule-backoffice/internal/domain/entity"

type IUserRepository interface {
	Create(user *entity.User) error
	FindByID(id int) (*entity.User, error)
	FindAll() (*entity.User, error)
}
