package repository

import (
	"github.com/br4tech/schedule-backoffice/internal/domain/entity"
	"github.com/br4tech/schedule-backoffice/internal/infra/database"
	"github.com/br4tech/schedule-backoffice/internal/infra/database/model"
	interfaces "github.com/br4tech/schedule-backoffice/internal/infra/repository/interface"
)

type UserRepository struct {
	adapter database.IGormAdapter
}

func NewUserRepository(adapter database.IGormAdapter) interfaces.IUserRepository {
	return &UserRepository{adapter: adapter}
}

func (r *UserRepository) Create(user *entity.User) error {
	userModel := model.ConvertUserToModel(user)
	return r.adapter.Create(userModel).Error
}

func (r *UserRepository) FindAll() (*entity.User, error) {

	var userModel model.User
	if err := r.adapter.FindAll(userModel).Error; err != nil {
		return nil, err
	}
	userDomain := model.ConvertUserToDomain(userModel)
	return userDomain, nil
}

func (r *UserRepository) FindByID(id int) (*entity.User, error) {
	user := &entity.User{}
	var userModel model.User
	if err := r.adapter.FindByID(userModel, id).Error; err != nil {
		return nil, err
	}
	return user, nil
}
