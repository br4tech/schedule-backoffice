package services

import (
	"github.com/br4tech/schedule-backoffice/internal/domain/entities"
	repositories "github.com/br4tech/schedule-backoffice/internal/repositories/interfaces"
)


type UserService struct {
	userRepository repositories.IUserRepository
}

func NewUserService(userRepository repositories.IUserRepository) *UserService {
   return &UserService{
       userRepository: userRepository,
   }
}

func (s *UserService) CreateUser(user *entities.User) error {
   return s.CreateUser(user)
}

func (s *UserService) GetUserByID(id uint) (*entities.User, error) {
   return s.GetUserByID(id)
}