package repositories

import (
	"github.com/br4tech/schedule-backoffice/internal/domain/entities"
	"github.com/jinzhu/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func(r *UserRepository) Create(user *entities.User) error {
	return r.db.Create(user).Error
}

func(r *UserRepository) FindByID(id int)(*entities.User, error){
	user := &entities.User{}
	if err := r.db.Where("id=?",id).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}