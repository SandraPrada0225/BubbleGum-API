package mocks

import (
	"bubblegum-api/internal/domain/dto/query"

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

func (mock *MockDulceProvider) GetDulcesListByCarritoID(id uint64) ([]uint64, error) {
	args := mock.Called(id)
	response := args.Get(0)
	err := args.Error(1)

	if response != nil {
		return response.([]uint64), err
	}
	return []uint64{}, err
}
