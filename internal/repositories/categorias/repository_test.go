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
	QuerySelectAll = "SELECT * FROM `Categorias`"
)

func TestGetAllOK(t *testing.T) {
	initialize()

	categorias := GetCategorias()
	t.Log(QuerySelectAll)
	mockDB.ExpectQuery(QuerySelectAll).
		WillReturnRows(sqlmock.NewRows([]string{"id", "nombre"}).
			AddRow(categorias[0].ID, categorias[0].Nombre).
			AddRow(categorias[1].ID, categorias[1].Nombre))

	categoriasRecibidas, err := repository.GetAll()

	assert.NoError(t, err)
	assert.Equal(t, categorias, categoriasRecibidas)
}

func TestByCodeInternalServerError(t *testing.T) {
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

func GetCategorias() (categorias []entities.Categoria) {
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
