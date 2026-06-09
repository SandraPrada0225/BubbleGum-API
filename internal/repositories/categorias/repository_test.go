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
	QuerySelectAll    = "SELECT * FROM `Categorias`"
)

func TestGetByCodeOK(t *testing.T) {
	initialize()

	mockCategoria := getMockCategorias()

	t.Log(QuerySelectByCode)
	mockDB.ExpectQuery(QuerySelectByCode).WithArgs(1).WillReturnRows(
		sqlmock.NewRows([]string{"id", "nombre"}).
			AddRow(mockCategoria[0].ID, mockCategoria[0].Nombre).
			AddRow(mockCategoria[1].ID, mockCategoria[1].Nombre))

	response, err := repository.GetCategoriasByDulceID(1)

	assert.NoError(t, err)
	assert.Equal(t, mockCategoria, response)
}

func TestGetAllOK(t *testing.T) {
	initialize()

	categorias := getMockCategorias()

	t.Log(QuerySelectAll)
	mockDB.ExpectQuery(QuerySelectAll).WillReturnRows(
		sqlmock.NewRows([]string{"id", "nombre"}).
			AddRow(categorias[0].ID, categorias[0].Nombre).
			AddRow(categorias[1].ID, categorias[1].Nombre))

	categoriasRecibidas, err := repository.GetAll()

	assert.NoError(t, err)
	assert.Equal(t, categorias, categoriasRecibidas)
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

func TestGetAllInternalServerError(t *testing.T) {
	initialize()
	mockDB.ExpectQuery(QuerySelectAll).WillReturnError(gorm.ErrInvalidData)

	categoriasRecibidas, err := repository.GetAll()

	typeErr := reflect.TypeOf(err).String()

	assert.Error(t, err)
	assert.Equal(t, "database.InternalServerError", typeErr)
	assert.Empty(t, categoriasRecibidas)
}

func initialize() {
	mockDB, DB = dbmocks.NewDB()
	mockDB.MatchExpectationsInOrder(false)
	repository = Repository{
		DB: DB.Debug(),
	}
}

func getMockCategorias() (categorias []entities.Categoria) {
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
