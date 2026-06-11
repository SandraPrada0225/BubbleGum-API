package mocks

import (
	"bubblegum-api/internal/domain/entities"

	"github.com/stretchr/testify/mock"
)

type MockUsuarioProvider struct {
	mock.Mock
}

func (mock *MockUsuarioProvider) GetByID(usuarioID uint64) (entities.Usuario, error) {
	args := mock.Called(usuarioID)
	response := args.Get(0)
	err := args.Error(1)

	if err != nil {
		return entities.Usuario{}, err
	}
	return response.(entities.Usuario), nil
}

func (mock *MockUsuarioProvider) Save(usuario *entities.Usuario) error {
	args := mock.Called(usuario)
	err := args.Error(0)
	return err
}
