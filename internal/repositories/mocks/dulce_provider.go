package mocks

import (
	"bubblegum-api/internal/domain/entities"

	"github.com/stretchr/testify/mock"
)

type MockDulceProvider struct {
	mock.Mock
}

func (mock *MockDulceProvider) GetByCode(codigo string) (entities.Dulce, error) {
	args := mock.Called(codigo)
	response := args.Get(0)
	err := args.Error(1)

	if response != nil {
		return response.(entities.Dulce), err
	}
	return entities.Dulce{}, err
}
