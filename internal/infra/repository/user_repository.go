package repository

import (
	"github.com/br4tech/schedule-backoffice/internal/domain/entity"
	"github.com/br4tech/schedule-backoffice/internal/infra/database"
	repository "github.com/br4tech/schedule-backoffice/internal/infra/repository/interfaces"
)

type UserRepository struct {
	adapter *database.GormAdapter
}

func NewUserRepository(adapter *database.GormAdapter) repository.IUserRepository {
	return &UserRepository{adapter: adapter}
}

func(r *UserRepository) Create(user *entity.User) error {
	return r.db.Create(user).Error
}

func(r *UserRepository) FindByID(id int)(*entity.User, error){
	user := &entity.User{}
	if err := r.db.Where("id=?",id).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}