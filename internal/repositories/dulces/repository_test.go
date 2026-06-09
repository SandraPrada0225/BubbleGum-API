package dulces

import (
	dbmocks "bubblegum-api/internal/app/config/database/mocks"
	"bubblegum-api/internal/domain/dto/responses"
	"bubblegum-api/internal/domain/entities"
	"reflect"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var (
	repository Repository
	mockDB     sqlmock.Sqlmock
	DB         *gorm.DB
)

const (
	MockDulceID                      = uint64(132423)
	MockCarritoID                    = uint64(2321)
	QuerySelectDulcesListByCarritoID = "SELECT * FROM `carritos_dulces` WHERE carrito_id=?"
	QuerySelectByCode                = "Call GetDetalleDulceByCode(?)"
	QuerySelectByID                  = "Call GetDetalleDulceByID(?)"
	QueryGetByID                     = "SELECT * FROM `dulces` WHERE id = ? LIMIT ?"
)

func TestGetByCodeOK(t *testing.T) {
	initialize()

	dulce := getResponse()

	t.Log(QuerySelectByCode)
	mockDB.ExpectQuery(QuerySelectByCode).WithArgs(dulce.Codigo).
		WillReturnRows(
			sqlmock.NewRows([]string{
				"id",
				"nombre",
				"presentacion_id",
				"presentacion_nombre",
				"descripcion",
				"imagen",
				"disponibles",
				"precio_unidad",
				"peso",
				"marca_id",
				"marca_nombre",
				"codigo"}).AddRow(
				dulce.ID,
				dulce.Nombre,
				dulce.Presentacion.ID,
				dulce.Presentacion.Nombre,
				dulce.Descripcion,
				dulce.Imagen,
				dulce.Disponibles,
				dulce.PrecioUnidad,
				dulce.Peso,
				dulce.Marca.ID,
				dulce.Marca.Nombre,
				dulce.Codigo,
			),
		)
	dulceRecibido, err := repository.GetByCode(dulce.Codigo)
	assert.NoError(t, err)
	assert.Equal(t, dulce, dulceRecibido)
}

func TestGetDetailByIDOK(t *testing.T) {
	initialize()

	dulce := getResponse()

	t.Log(QuerySelectByID)
	mockDB.ExpectQuery(QuerySelectByID).WithArgs(dulce.ID).
		WillReturnRows(
			sqlmock.NewRows([]string{
				"id",
				"nombre",
				"presentacion_id",
				"presentacion_nombre",
				"descripcion",
				"imagen",
				"disponibles",
				"precio_unidad",
				"peso",
				"marca_id",
				"marca_nombre",
				"codigo"}).AddRow(
				dulce.ID,
				dulce.Nombre,
				dulce.Presentacion.ID,
				dulce.Presentacion.Nombre,
				dulce.Descripcion,
				dulce.Imagen,
				dulce.Disponibles,
				dulce.PrecioUnidad,
				dulce.Peso,
				dulce.Marca.ID,
				dulce.Marca.Nombre,
				dulce.Codigo,
			),
		)
	dulceRecibido, err := repository.GetDetailByID(dulce.ID)
	assert.NoError(t, err)
	assert.Equal(t, dulce, dulceRecibido)
}

func TestByCodeErrorNotFound(t *testing.T) {
	initialize()

	mockDB.ExpectQuery(QuerySelectByCode).WithArgs("2").WillReturnError(gorm.ErrRecordNotFound)

	dulceRecibido, err := repository.GetByCode("2")

	typeErr := reflect.TypeOf(err).String()

	assert.Error(t, err)
	assert.Equal(t, "database.NotFoundError", typeErr)
	assert.Empty(t, dulceRecibido)
}

func TestByCodeInternalServerError(t *testing.T) {
	initialize()

	mockDB.ExpectQuery(QuerySelectByCode).WithArgs("2").WillReturnError(gorm.ErrInvalidData)

	dulceRecibido, err := repository.GetByCode("2")

	typeErr := reflect.TypeOf(err).String()

	assert.Error(t, err)
	assert.Equal(t, "database.InternalServerError", typeErr)
	assert.Empty(t, dulceRecibido)
}

func TestGetDetailByIDErrorNotFound(t *testing.T) {
	initialize()

	mockDB.ExpectQuery(QuerySelectByID).WithArgs(MockDulceID).WillReturnError(gorm.ErrRecordNotFound)

	dulceRecibido, err := repository.GetDetailByID(MockDulceID)

	typeErr := reflect.TypeOf(err).String()

	assert.Error(t, err)
	assert.Equal(t, "database.NotFoundError", typeErr)
	assert.Empty(t, dulceRecibido)
}

func TestGetDetailByIDInternalServerError(t *testing.T) {
	initialize()

	mockDB.ExpectQuery(QuerySelectByID).WithArgs(2).WillReturnError(gorm.ErrInvalidData)

	dulceRecibido, err := repository.GetDetailByID(2)

	typeErr := reflect.TypeOf(err).String()

	assert.Error(t, err)
	assert.Equal(t, "database.InternalServerError", typeErr)
	assert.Empty(t, dulceRecibido)
}
func initialize() {
	mockDB, DB = dbmocks.NewDB()
	mockDB.MatchExpectationsInOrder(false)
	repository = Repository{
		DB: DB.Debug(),
	}
}

func TestGetByIDOK(t *testing.T) {
	initialize()

	dulce := getMockDulce()

	t.Log(QueryGetByID)
	mockDB.ExpectQuery(QueryGetByID).WithArgs(dulce.ID, 1).
		WillReturnRows(
			sqlmock.NewRows([]string{
				"id",
				"nombre",
				"marca_id",
				"precio",
				"peso",
				"unidades",
				"presentacion_id",
				"descripcion",
				"imagen",
				"fecha_vencimiento",
				"fecha_expedicion",
				"disponibles",
				"codigo",
			}).AddRow(
				dulce.ID,
				dulce.Nombre,
				dulce.MarcaID,
				dulce.Precio,
				dulce.Peso,
				dulce.Unidades,
				dulce.PresentacionID,
				dulce.Descripcion,
				dulce.Imagen,
				dulce.FechaVencimiento,
				dulce.FechaExpedicion,
				dulce.Disponibles,
				dulce.Codigo,
			),
		)
	dulceRecibido, err := repository.GetByID(dulce.ID)
	assert.NoError(t, err)
	assert.Equal(t, dulce, dulceRecibido)
}

func TestGetByIDErrorNotFound(t *testing.T) {
	initialize()

	mockDB.ExpectQuery(QueryGetByID).WithArgs(2, 1).WillReturnError(gorm.ErrRecordNotFound)

	dulceRecibido, err := repository.GetByID(2)

	typeErr := reflect.TypeOf(err).String()

	assert.Error(t, err)
	assert.Equal(t, "database.NotFoundError", typeErr)
	assert.Empty(t, dulceRecibido)
}

func TestGetByIDInternalServerError(t *testing.T) {
	initialize()

	mockDB.ExpectQuery(QueryGetByID).WithArgs(2, 1).WillReturnError(gorm.ErrInvalidData)

	dulceRecibido, err := repository.GetDetailByID(2)

	typeErr := reflect.TypeOf(err).String()

	assert.Error(t, err)
	assert.Equal(t, "database.InternalServerError", typeErr)
	assert.Empty(t, dulceRecibido)
}

func TestGetDulcesListByCarritoIDWhenEveryThingWentSuccessfullyShouldReturnDulcesList(t *testing.T) {
	initialize()

	dulcesList := getMockDulcesInCarritoList()

	t.Log(QuerySelectDulcesListByCarritoID)
	mockDB.ExpectQuery(QuerySelectDulcesListByCarritoID).WithArgs(MockCarritoID).
		WillReturnRows(
			sqlmock.NewRows([]string{
				"id",
				"dulce_id",
				"carrito_id",
				"unidades",
				"subtotal",
			}).AddRow(
				dulcesList[0].ID,
				dulcesList[0].DulceID,
				dulcesList[0].CarritoID,
				dulcesList[0].Unidades,
				dulcesList[0].Subtotal,
			).AddRow(
				dulcesList[1].ID,
				dulcesList[1].DulceID,
				dulcesList[1].CarritoID,
				dulcesList[1].Unidades,
				dulcesList[1].Subtotal,
			),
		)
	dulcesListRecibidos, err := repository.GetDulcesListByCarritoID(MockCarritoID)
	assert.NoError(t, err)
	assert.Equal(t, dulcesList, dulcesListRecibidos)
}

func TestGetDulcesListByCarritoIDWhenSomethingWentWrongShouldReturnInternalError(t *testing.T) {
	initialize()

	mockDB.ExpectQuery(QuerySelectDulcesListByCarritoID).WithArgs(MockCarritoID).WillReturnError(gorm.ErrInvalidData)

	dulcesListResponse, err := repository.GetDulcesListByCarritoID(MockCarritoID)

	assert.Error(t, err)
	assert.Empty(t, dulcesListResponse)
	assert.NoError(t, mockDB.ExpectationsWereMet())
}

func getResponse() (response responses.DetalleDulce) {
	response = responses.DetalleDulce{
		ID:     MockDulceID,
		Nombre: "Chocolatina",
		Presentacion: entities.Presentacion{
			ID:     1,
			Nombre: "Empaque",
		},
		Descripcion:  "Deliciosa chocolatina que se derrite en tu boca",
		Imagen:       "imagen",
		Disponibles:  100,
		PrecioUnidad: 1000,
		Peso:         40,
		Marca: entities.Marca{
			ID:     1,
			Nombre: "Jet",
		},
		Codigo: "2",
	}
	return
}

func getMockDulcesInCarritoList() []entities.CarritoDulce {
	return []entities.CarritoDulce{
		{
			ID:        1,
			CarritoID: 1,
			DulceID:   1,
			Unidades:  2,
			Subtotal:  2000,
		},
		{
			ID:        2,
			CarritoID: 2,
			DulceID:   1,
			Unidades:  1,
			Subtotal:  1000,
		},
	}
}

func getMockDulce() (dulce entities.Dulce) {
	dulce = entities.Dulce{
		ID:               1,
		Nombre:           "Gomas clasicas",
		MarcaID:          6,
		Precio:           2950.000,
		Peso:             80,
		Unidades:         5,
		PresentacionID:   4,
		Descripcion:      "Gomas clasicas con sabores surtidos",
		FechaVencimiento: time.Date(2023, time.August, 24, 0, 0, 0, 0, time.Local),
		FechaExpedicion:  time.Date(2023, time.July, 24, 0, 0, 0, 0, time.Local),
		Disponibles:      100,
		Codigo:           "1A",
	}
	return
}
