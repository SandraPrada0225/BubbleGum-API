package getpurchaselistbyuserid

import (
	"bubblegum-api/internal/domain/dto/responses"
	"bubblegum-api/internal/domain/errors/database"
	"bubblegum-api/internal/repositories/mocks"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var mockVentasProvider *mocks.MockVentaProvider

const (
	mockUserID uint64 = 5
)

func TestWhenIsSuccesfullShouldReturnListAndNoError(t *testing.T) {
	useCase := initialize()
	mockedPurchaseList := getMockedPurchaseList()

	mockVentasProvider.On("GetListByUserID", mockUserID).Return(mockedPurchaseList, nil)

	response, err := useCase.Execute(mockUserID)

	assert.NoError(t, err)
	assert.Equal(t, mockedPurchaseList, response)
	mockVentasProvider.AssertNumberOfCalls(t, "GetListByUserID", 1)
}

func TestWhenGetListWentWrongShouldReturnInternalServerError(t *testing.T) {
	useCase := initialize()
	mockVentasProvider.On("GetListByUserID", mockUserID).Return(responses.GetPurchaseList{}, database.NewInterlServerError("error"))

	response, err := useCase.Execute(mockUserID)
	typeErr := reflect.TypeOf(err).String()

	assert.Error(t, err)
	assert.Equal(t, "database.InternalServerError", typeErr)
	assert.Empty(t, response)
	mockVentasProvider.AssertNumberOfCalls(t, "GetListByUserID", 1)
}

func initialize() Implementation {
	mockVentasProvider = new(mocks.MockVentaProvider)
	return Implementation{
		VentasProvider: mockVentasProvider,
	}
}

func getMockedPurchaseList() responses.GetPurchaseList {
	fecha := time.Date(2023, 10, 17, 0, 0, 0, 0, time.Local)
	return responses.GetPurchaseList{
		PurchaseList: []responses.Purchase{
			{
				ID:              1,
				Fecha:           fecha,
				MedioDePagoID:   1,
				MedioDePago:     "contraentrega",
				PrecioTotal:     100,
				Subtotal:        97,
				Descuento:       2,
				Envio:           5,
				EstadoCarritoID: 1,
				EstadoCarrito:   "comprador",
			},
			{
				ID:              2,
				Fecha:           fecha,
				MedioDePagoID:   1,
				MedioDePago:     "credito",
				PrecioTotal:     120,
				Subtotal:        100,
				Descuento:       0,
				Envio:           20,
				EstadoCarritoID: 1,
				EstadoCarrito:   "comprador",
			},
		},
	}
}
