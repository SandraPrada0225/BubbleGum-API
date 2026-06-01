package presentaciones

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
	QuerySelectAll = "SELECT * FROM `Presentaciones`"
)

func TestGetAllOK(t *testing.T) {
	initialize()

	presentaciones := GetPresentaciones()
	t.Log(QuerySelectAll)
	mockDB.ExpectQuery(QuerySelectAll).
		WillReturnRows(sqlmock.NewRows([]string{"id", "nombre"}).
			AddRow(presentaciones[0].ID, presentaciones[0].Nombre).
			AddRow(presentaciones[1].ID, presentaciones[1].Nombre))

	presentacionesRecibidas, err := repository.GetAll()

	assert.NoError(t, err)
	assert.Equal(t, presentaciones, presentacionesRecibidas)
}

func TestByCodeInternalServerError(t *testing.T) {
	initialize()
	mockDB.ExpectQuery(QuerySelectAll).WillReturnError(gorm.ErrInvalidData)

	presentacionesRecibidas, err := repository.GetAll()

	typeErr := reflect.TypeOf(err).String()

	assert.Error(t, err)
	assert.Equal(t, "database.InternalServerError", typeErr)
	assert.Empty(t, presentacionesRecibidas)
}

func initialize() {
	mockDB, DB = dbmocks.NewDB()
	mockDB.MatchExpectationsInOrder(false)
	repository = Repository{
		DB: DB.Debug(),
	}
}

func GetPresentaciones() (presentaciones []entities.Presentacion) {
	presentaciones = []entities.Presentacion{
		{
			ID:     1,
			Nombre: "Caja",
		},
		{
			ID:     2,
			Nombre: "Bolsa",
		},
	}

	return
}
