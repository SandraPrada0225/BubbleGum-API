package categorias

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
	QuerySelectByCode = "Call GetCategoriasByDulceID(?)"
)

func TestGetByCodeOK(t *testing.T) {
	initialize()

	mockCategoria := GetMockCategorias()

	t.Log(QuerySelectByCode)
	mockDB.ExpectQuery(QuerySelectByCode).WithArgs(1).WillReturnRows(
		sqlmock.NewRows([]string{"id", "nombre"}).
			AddRow(mockCategoria[0].ID, mockCategoria[0].Nombre).
			AddRow(mockCategoria[1].ID, mockCategoria[1].Nombre))

	response, err := repository.GetCategoriasByDulceID(1)

	assert.NoError(t, err)
	assert.Equal(t, mockCategoria, response)
}

func TestByCodeInternalServerError(t *testing.T) {
	initialize()
	mockDB.ExpectQuery(QuerySelectByCode).WithArgs(1).WillReturnError(gorm.ErrInvalidData)

	dulceRecibido, err := repository.GetCategoriasByDulceID(1)

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

func GetMockCategorias() (categorias []entities.Categoria) {
	categorias = []entities.Categoria{
		{
			ID:     1,
			Nombre: "Gomitas",
		},
		{
			ID:     2,
			Nombre: "Chupetes",
		},
	}
	return
}
