package service

import (
	"github.com/br4tech/schedule-backoffice/internal/domain/entity"
	service "github.com/br4tech/schedule-backoffice/internal/domain/service/interface"
	repository "github.com/br4tech/schedule-backoffice/internal/infra/repository/interfaces"
)


type UserService struct {
	userRepository repository.IUserRepository
}

func NewUserService(userRepository repository.IUserRepository) service.IUserService {
   return &UserService{ userRepository: userRepository }
}

func (s *UserService) CreateUser(user *entity.User) error {
   return s.CreateUser(user)
}

func (s *UserService) GetUserByID(id int) (*entity.User, error) {
   return s.GetUserByID(id)
}