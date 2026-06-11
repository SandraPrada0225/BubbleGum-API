package usuarios

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
	updateQuery = "UPDATE `usuarios` SET `nombre`=?,`apellido`=?,`password`=?,`correo`=?,`carrito_actual_id`=? WHERE `id` = ?"
	selectQuery = "SELECT * FROM `usuarios` WHERE id = ? LIMIT ?"
)

func TestGetByIDWhenIsSuccesfullReturnUsuario(t *testing.T) {
	initialize()

	mockUsuario := getMockUsuario()

	t.Log(selectQuery)
	mockDB.ExpectQuery(selectQuery).WithArgs(mockUsuario.ID, 1).WillReturnRows(
		sqlmock.NewRows([]string{
			"id",
			"nombre",
			"apellido",
			"correo",
			"password",
			"carrito_actual_id"}).
			AddRow(
				mockUsuario.ID,
				mockUsuario.Nombre,
				mockUsuario.Apellido,
				mockUsuario.Correo,
				mockUsuario.Password,
				mockUsuario.CarritoActualID))

	usuario, err := repository.GetByID(mockUsuario.ID)

	assert.NoError(t, err)
	assert.Equal(t, mockUsuario, usuario)
	assert.NoError(t, mockDB.ExpectationsWereMet())
}

func TestGetByIDWhenUsuarioDoesNotExistsShouldReturnNotfoundError(t *testing.T) {
	initialize()

	mockUsuario := getMockUsuario()

	mockDB.ExpectQuery(selectQuery).WithArgs(mockUsuario.ID, 1).WillReturnError(gorm.ErrRecordNotFound)

	usuario, err := repository.GetByID(mockUsuario.ID)

	typeErr := reflect.TypeOf(err).String()

	assert.Error(t, err)
	assert.Equal(t, "database.NotFoundError", typeErr)
	assert.Empty(t, usuario)
	assert.NoError(t, mockDB.ExpectationsWereMet())
}

func TestGetByIDWhenWrongShouldReturnInternalError(t *testing.T) {
	initialize()
	mockUsuario := getMockUsuario()

	mockDB.ExpectQuery(selectQuery).WithArgs(mockUsuario.ID, 1).WillReturnError(gorm.ErrInvalidData)

	usuario, err := repository.GetByID(mockUsuario.ID)

	typeErr := reflect.TypeOf(err).String()

	assert.Error(t, err)
	assert.Equal(t, "database.InternalServerError", typeErr)
	assert.Empty(t, usuario)
	assert.NoError(t, mockDB.ExpectationsWereMet())
}

func TestSaveWhenIsSuccesfullReturnNoError(t *testing.T) {
	initialize()
	mockDB.ExpectBegin()

	t.Log(updateQuery)
	mockDB.ExpectExec(updateQuery).WillReturnResult(sqlmock.NewResult(1, 1))
	mockDB.ExpectCommit()
	mockUsuario := getMockUsuario()

	err := repository.Save(&mockUsuario)

	assert.NoError(t, err)
	assert.NoError(t, mockDB.ExpectationsWereMet())
}

func TestSaveWhenWentWrongShouldReturnInternalError(t *testing.T) {
	initialize()

	mockDB.ExpectBegin()
	mockDB.ExpectExec(updateQuery).WillReturnError(gorm.ErrInvalidData)
	mockDB.ExpectRollback()

	mockUsuario := getMockUsuario()
	err := repository.Save(&mockUsuario)
	typeErr := reflect.TypeOf(err).String()

	assert.Error(t, err)
	assert.Equal(t, "database.InternalServerError", typeErr)
}

func TestSaveWhenCarBelongsToOtherUserShouldReturnConflictError(t *testing.T) {
	initialize()
	mockDB.ExpectBegin()
	mockDB.ExpectExec(updateQuery).WillReturnError(gorm.ErrDuplicatedKey)
	mockDB.ExpectRollback()
	mockUsuario := getMockUsuario()

	err := repository.Save(&mockUsuario)

	assert.Error(t, err)
	assert.Equal(t, "database.ConflictError", reflect.TypeOf(err).String())
}

func initialize() {
	mockDB, DB = dbmocks.NewDB()
	mockDB.MatchExpectationsInOrder(false)
	repository = Repository{
		DB: DB.Debug(),
	}
}

func getMockUsuario() entities.Usuario {
	return entities.Usuario{
		ID:              1,
		Nombre:          "Milena",
		Apellido:        "Prada",
		CarritoActualID: 2,
		Password:        "Sandra123-",
	}
}
