package services

import (
	"github.com/br4tech/schedule-backoffice/internal/domain/entities"
	services "github.com/br4tech/schedule-backoffice/internal/domain/services/interface"
	repositories "github.com/br4tech/schedule-backoffice/internal/infra/repositories/interfaces"
)


type UserService struct {
	userRepository repositories.IUserRepository
}

func NewUserService(userRepository repositories.IUserRepository) services.IUserService {
   return &UserService{
       userRepository: userRepository,
   }
}

func (s *UserService) CreateUser(user *entities.User) error {
   return s.CreateUser(user)
}

func (s *UserService) FindByID(id int) (*entities.User, error) {
   return s.FindByID(id)
}