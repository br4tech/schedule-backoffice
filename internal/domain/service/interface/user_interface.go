package service

import "github.com/br4tech/schedule-backoffice/internal/domain/entity"

type IUserService interface {
	CreateUser(user *entity.User) error
	GetUserByID(id int) (*entity.User, error)
}
