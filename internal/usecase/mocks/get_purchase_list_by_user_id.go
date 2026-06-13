package mocks

import (
	"bubblegum-api/internal/domain/dto/responses"

	"github.com/stretchr/testify/mock"
)

type MockGetPurchaseListByUserID struct {
	mock.Mock
}

func (m *MockGetPurchaseListByUserID) Execute(userID uint64) (responses.GetPurchaseList, error) {
	args := m.Called(userID)
	response := args.Get(0)
	err := args.Error(1)
	if err != nil {
		return responses.GetPurchaseList{}, err
	}
	return response.(responses.GetPurchaseList), err
}
