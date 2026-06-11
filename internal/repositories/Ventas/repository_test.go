package ventas

import (
	dbmocks "bubblegum-api/internal/app/config/database/mocks"
	"bubblegum-api/internal/domain/entities"
	"reflect"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

const (
	createQuery = "INSERT INTO `ventas` (`medio_de_pago_id`,`carrito_id`,`comprador_id`,`created_at`) VALUES (?,?,?,?)"
)

var (
	repository Repository
	mockDB     sqlmock.Sqlmock
	DB         *gorm.DB
)

func TestWhenCreatedWasSuccesfullAndShouldReturnNoError(t *testing.T) {
	initialize()

	ventaToCreate := getMockedVentaToCreate()

	t.Log(createQuery)
	mockDB.ExpectBegin()
	mockDB.ExpectExec(createQuery).WillReturnResult(sqlmock.NewResult(1, 1))
	mockDB.ExpectCommit()

	err := repository.Create(&ventaToCreate)

	assert.NoError(t, err)
	assert.NoError(t, mockDB.ExpectationsWereMet())
}

func TestWhenCarritoWasNotFoundShouldReturnNotfoundError(t *testing.T) {
	initialize()

	ventaToCreate := getMockedVentaToCreate()

	mockDB.ExpectBegin()
	mockDB.ExpectExec(createQuery).WillReturnError(gorm.ErrForeignKeyViolated)
	mockDB.ExpectRollback()

	err := repository.Create(&ventaToCreate)

	typeErr := reflect.TypeOf(err).String()

	assert.Error(t, err)
	assert.Equal(t, "database.NotFoundError", typeErr)
	assert.NoError(t, mockDB.ExpectationsWereMet())
}

func TestWhenCreateWhenWrongShouldReturnInternalError(t *testing.T) {
	initialize()

	ventaToCreate := getMockedVentaToCreate()

	mockDB.ExpectBegin()
	mockDB.ExpectExec(createQuery).WillReturnError(gorm.ErrInvalidDB)
	mockDB.ExpectRollback()

	err := repository.Create(&ventaToCreate)

	typeErr := reflect.TypeOf(err).String()

	assert.Error(t, err)
	assert.Equal(t, "database.InternalServerError", typeErr)
	assert.NoError(t, mockDB.ExpectationsWereMet())
}

func TestWhenCarritoWasPurchasedShouldReturnConflictError(t *testing.T) {
	initialize()

	ventaToCreate := getMockedVentaToCreate()

	mockDB.ExpectBegin()
	mockDB.ExpectExec(createQuery).WillReturnError(gorm.ErrDuplicatedKey)
	mockDB.ExpectRollback()

	err := repository.Create(&ventaToCreate)
	typeErr := reflect.TypeOf(err).String()

	assert.Error(t, err)
	assert.Equal(t, "database.ConflictError", typeErr)
	assert.NoError(t, mockDB.ExpectationsWereMet())
}

func initialize() {
	mockDB, DB = dbmocks.NewDB()
	mockDB.MatchExpectationsInOrder(false)
	repository = Repository{
		DB: DB.Debug(),
	}
}

func getMockedVentaToCreate() entities.Venta {
	fecha := time.Date(2023, 10, 17, 0, 0, 0, 0, time.Local)
	return entities.Venta{
		CarritoID:     1,
		CreatedAt:     fecha,
		MedioDePagoID: 1,
	}
}
