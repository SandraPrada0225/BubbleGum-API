package carritos

import (
	dbmocks "bubblegum-api/internal/app/config/database/mocks"
	"bubblegum-api/internal/domain/entities"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

const (
	mockCarritoID          = uint64(132423)
	QuerySelectCarritoByID = "SELECT * FROM `carritos` WHERE id = ? LIMIT ?"
)

var (
	repository Repository
	mockDB     sqlmock.Sqlmock
	DB         *gorm.DB
)

func TestGetCarritoByCarritoIDOK(t *testing.T) {
	initialize()

	carrito := GetResponse()
	t.Log(QuerySelectCarritoByID)
	mockDB.ExpectQuery(QuerySelectCarritoByID).WithArgs(carrito.ID, 1).
		WillReturnRows(
			sqlmock.NewRows([]string{
				"id",
				"precio_total",
				"subtotal",
				"envio",
				"descuento",
			}).AddRow(
				carrito.ID,
				carrito.PrecioTotal,
				carrito.Subtotal,
				carrito.Envio,
				carrito.Descuento,
			),
		)
	carritoRecibido, err := repository.GetCarritoByCarritoID(carrito.ID)
	assert.NoError(t, err)
	assert.Equal(t, carrito, carritoRecibido)
}

func TestGetCarritoByCarritoIDErrorNotFound(t *testing.T) {
	initialize()

	mockDB.ExpectQuery(QuerySelectCarritoByID).WithArgs(mockCarritoID, 1).WillReturnError(gorm.ErrRecordNotFound)

	carritoRecibido, err := repository.GetCarritoByCarritoID(mockCarritoID)

	typeErr := reflect.TypeOf(err).String()

	assert.Error(t, err)
	assert.Equal(t, "database.NotFoundError", typeErr)
	assert.Empty(t, carritoRecibido)
}

func TestGetCarritoByCarritoIDInternalServerError(t *testing.T) {
	initialize()

	mockDB.ExpectQuery(QuerySelectCarritoByID).WithArgs(mockCarritoID, 1).WillReturnError(gorm.ErrInvalidData)

	carritoRecibido, err := repository.GetCarritoByCarritoID(mockCarritoID)

	typeErr := reflect.TypeOf(err).String()

	assert.Error(t, err)
	assert.Equal(t, "database.InternalServerError", typeErr)
	assert.Empty(t, carritoRecibido)
}

func initialize() {
	mockDB, DB = dbmocks.NewDB()
	mockDB.MatchExpectationsInOrder(false)
	repository = Repository{
		DB: DB.Debug(),
	}
}

func GetResponse() (response entities.Carrito) {
	response = entities.Carrito{
		ID:          mockCarritoID,
		PrecioTotal: 1000,
		Subtotal:    995,
		Envio:       5,
		Descuento:   0,
	}
	return
}
