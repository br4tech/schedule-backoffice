package repository_test

import (
	"testing"

	"github.com/br4tech/schedule-backoffice/internal/domain/entity"
	"github.com/br4tech/schedule-backoffice/internal/infra/database"
	"github.com/br4tech/schedule-backoffice/internal/infra/repository"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestUserRepository_Create(t *testing.T) {
	dsn := "root:admin@tcp(localhost:3306)/schedule_app?charset=utf8mb4&parseTime=True"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	assert.NoError(t, err, "Erro ao abrir conexão com o banco de dados")

	adapter := database.NewGormAdapter(db)
	// defer adapter.Close()

	userRepository := repository.NewUserRepository(adapter)

	user := &entity.User{
		Name:     "Guilherme",
		LastName: "Silva",
		Email:    "guilherme.silva@gmail.com",
	}

	err = userRepository.Create(user)

	assert.NoError(t, err, "Erro ao criar usuario")
}
