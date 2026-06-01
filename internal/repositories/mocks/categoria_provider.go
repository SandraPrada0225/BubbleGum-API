package mocks

import (
	"bubblegum-api/internal/domain/entities"

	"github.com/stretchr/testify/mock"
)

type MockCategoriaProvider struct {
	mock.Mock
}

func (mock *MockCategoriaProvider) GetAll() ([]entities.Categoria, error) {
	args := mock.Called()
	response := args.Get(0)
	err := args.Error(1)

	if response != nil {
		return response.([]entities.Categoria), err
	}
	return []entities.Categoria{}, err
}
