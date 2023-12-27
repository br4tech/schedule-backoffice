package service

import (
	"errors"
	"testing"

	"github.com/br4tech/schedule-backoffice/internal/domain/entity"
)

type MockUserRepository struct{}

func (m *MockUserRepository) GetUserByID(id int) (*entity.User, error) {
	if id == 1 {
		return &entity.User{ID: 1, Name: "Joao"}, nil
	}
	return nil, errors.New("Usuario nao encontrado")
}

func TestUserService_GetUserByID(t *testing.T) {
	userService := NewUserService(&MockUserRepository{})

	user, err := userService.GetUserByID(1)
	if err != nil {
		t.Errorf("Erro inesperado %v", err)
	}

	if user == nil {
		t.Error("Usuario nao encontrado")
	}

	if user.Name != "Joao" {
		t.Errorf("Nome do usuario incorreto.Esperado 'Joao', mas obteve '%s'", user.Name)
	}

	user, err = userService.GetUserByID(2)
	if err == nil {
		t.Error("Esperava um erro, mas não ocorreu")
	}
	if user != nil {
		t.Error("O usuário deveria ser nulo")
	}

}
