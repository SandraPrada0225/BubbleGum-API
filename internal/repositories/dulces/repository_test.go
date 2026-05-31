package dulces

import (
	dbmocks "bubblegum-api/internal/app/config/database/mocks"
	"bubblegum-api/internal/domain/entities"
	"reflect"
	"testing"

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
	QuerySelectByCode = "SELECT * FROM `dulces` WHERE codigo = ? ORDER BY `dulces`.`id` LIMIT ?"
)

func TestGetByCodeOK(t *testing.T) {
	initialize()

	dulce := GetDulce()
	t.Log(QuerySelectByCode)
	mockDB.ExpectQuery(QuerySelectByCode).WithArgs(dulce.Codigo, 1).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "nombre", "presentacion_id", "descripcion", "imagen",
				"disponibles", "precio", "peso", "marca_id", "codigo"}).
				AddRow(
					dulce.ID,
					dulce.Nombre,
					dulce.PresentacionID,
					dulce.Descripcion,
					dulce.Imagen,
					dulce.Disponibles,
					dulce.Precio,
					dulce.Peso,
					dulce.MarcaID,
					dulce.Codigo,
				),
		)
	dulceRecibido, err := repository.GetByCode(dulce.Codigo)
	assert.NoError(t, err)
	assert.Equal(t, dulce, dulceRecibido)
}

func TestByCodeErrorNotFound(t *testing.T) {
	initialize()

	mockDB.ExpectQuery(QuerySelectByCode).WithArgs("2", 1).WillReturnError(gorm.ErrRecordNotFound)

	dulceRecibido, err := repository.GetByCode("2")

	typeErr := reflect.TypeOf(err).String()

	assert.Error(t, err)
	assert.Equal(t, "database.NotFoundError", typeErr)
	assert.Empty(t, dulceRecibido)
}

func TestByCodeInternalServerError(t *testing.T) {
	initialize()

	mockDB.ExpectQuery(QuerySelectByCode).WithArgs("2", 1).WillReturnError(gorm.ErrInvalidData)

	dulceRecibido, err := repository.GetByCode("2")

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

func GetDulce() (dulce entities.Dulce) {
	dulce = entities.Dulce{
		ID:             2,
		Nombre:         "Chocolatina",
		PresentacionID: 1,
		Descripcion:    "Deliciosa chocolatina que se derrite en tu boca",
		Imagen:         "imagen",
		Disponibles:    100,
		Precio:         1000,
		Peso:           40,
		MarcaID:        1,
		Codigo:         "2",
	}
	return
}
