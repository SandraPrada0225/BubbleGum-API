package mocks

import (
	"bubblegum-api/internal/domain/dto/query"
	"bubblegum-api/internal/domain/entities"

	"github.com/stretchr/testify/mock"
)

type MockDulceProvider struct {
	mock.Mock
}

func (mock *MockDulceProvider) GetByCode(codigo string) (query.DetalleDulce, error) {
	args := mock.Called(codigo)
	response := args.Get(0)
	err := args.Error(1)

	if response != nil {
		return response.(query.DetalleDulce), err
	}
	return query.DetalleDulce{}, err
}

func (mock *MockDulceProvider) GetDetailByID(id uint64) (query.DetalleDulce, error) {
	args := mock.Called(id)
	response := args.Get(0)
	err := args.Error(1)

	if response != nil {
		return response.(query.DetalleDulce), err
	}
	return query.DetalleDulce{}, err
}

func (mock *MockDulceProvider) GetDulcesListByCarritoID(id uint64) ([]entities.CarritoDulce, error) {
	args := mock.Called(id)
	response := args.Get(0)
	err := args.Error(1)

	if response != nil {
		return response.([]entities.CarritoDulce), err
	}
	return []entities.CarritoDulce{}, err
}
