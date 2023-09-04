package services

import "github.com/br4tech/schedule-backoffice/internal/domain/entities"

type IUserService interface {
	CreateUser(user *entities.User) error
	FindByID(id int) (*entities.User, error)
}