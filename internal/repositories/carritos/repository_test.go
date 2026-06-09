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
	mockCarritoID                      = uint64(2)
	queryCreate                        = "INSERT INTO `carritos` (`precio_total`,`subtotal`,`envio`,`descuento`) VALUES (?,?,?,?)"
	queryUpdate                        = "UPDATE `carritos` SET `precio_total`=?,`subtotal`=?,`envio`=?,`descuento`=? WHERE `id` = ?"
	QuerySelectByID                    = "SELECT * FROM `carritos` WHERE id = ? LIMIT ?"
	QueryGetDulceByCarritoIDAndDulceID = "SELECT * FROM `carritos_dulces` WHERE carrito_id = ? AND dulce_id = ? ORDER BY `carritos_dulces`.`id` LIMIT ?"
	QueryUpdateDulceInCarrito          = "UPDATE `carritos_dulces` SET `carrito_id`=?,`dulce_id`=?,`unidades`=?,`subtotal`=? WHERE `id` = ?"
	QueryAddDulceInCarrito             = "INSERT INTO `carritos_dulces` (`carrito_id`,`dulce_id`,`unidades`,`subtotal`) VALUES (?,?,?,?)"
	QueryDeleteDulceInCarrito          = "DELETE FROM `carritos_dulces` WHERE `carritos_dulces`.`id` = ?"
)

var (
	repository Repository
	mockDB     sqlmock.Sqlmock
	DB         *gorm.DB
)

func TestGeByIDOK(t *testing.T) {
	initialize()

	carrito := getResponse()
	t.Log(QuerySelectByID)
	mockDB.ExpectQuery(QuerySelectByID).WithArgs(carrito.ID, 1).
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
	carritoRecibido, err := repository.GetByID(carrito.ID)
	assert.NoError(t, err)
	assert.Equal(t, carrito, carritoRecibido)
}

func TestGetDulcesByCarritoIDAndDulceIDOK(t *testing.T) {
	initialize()

	carritoDulce := getMockCarritoDulce()
	t.Log(QueryGetDulceByCarritoIDAndDulceID)

	mockDB.ExpectQuery(QueryGetDulceByCarritoIDAndDulceID).WithArgs(carritoDulce.CarritoID, carritoDulce.DulceID, 1).
		WillReturnRows(
			sqlmock.NewRows([]string{
				"id",
				"carrito_id",
				"dulce_id",
				"unidades",
				"subtotal",
			}).AddRow(
				carritoDulce.ID,
				carritoDulce.CarritoID,
				carritoDulce.DulceID,
				carritoDulce.Unidades,
				carritoDulce.Subtotal,
			),
		)
	carritoDulceRecibido, exists, err := repository.GetDulceByCarritoIDAndDulceID(carritoDulce.CarritoID, carritoDulce.DulceID)
	assert.NoError(t, err)
	assert.True(t, exists)
	assert.Equal(t, carritoDulce, carritoDulceRecibido)
	assert.NoError(t, mockDB.ExpectationsWereMet())
}

func TestGetDulceByCarritoIDAndDulceIDWhenCartDoesNotexistReturnFalse(t *testing.T) {
	initialize()

	mockDB.ExpectQuery(QueryGetDulceByCarritoIDAndDulceID).WithArgs(2, 2, 1).WillReturnError(gorm.ErrRecordNotFound)

	carritoDulceRecibido, exists, err := repository.GetDulceByCarritoIDAndDulceID(2, 2)

	assert.NoError(t, err)
	assert.False(t, exists)
	assert.Empty(t, carritoDulceRecibido)
	assert.NoError(t, mockDB.ExpectationsWereMet())
}

func TestGetDulceByCarritoIDAndDulceIDInternalServerError(t *testing.T) {
	initialize()

	mockDB.ExpectQuery(QueryGetDulceByCarritoIDAndDulceID).WithArgs(2, 2, 1).WillReturnError(gorm.ErrInvalidData)

	carritoDulceRecibido, exists, err := repository.GetDulceByCarritoIDAndDulceID(2, 2)

	typeErr := reflect.TypeOf(err).String()

	assert.Error(t, err)
	assert.False(t, exists)
	assert.Equal(t, "database.InternalServerError", typeErr)
	assert.Empty(t, carritoDulceRecibido)
	assert.NoError(t, mockDB.ExpectationsWereMet())
}

func TestUpdateDulceInCarritoOK(t *testing.T) {
	initialize()

	carritoDulce := getMockCarritoDulce()

	mockDB.ExpectBegin()
	mockDB.ExpectExec(QueryUpdateDulceInCarrito).WillReturnResult(sqlmock.NewResult(1, 1))
	mockDB.ExpectCommit()

	err := repository.AddDulceInCarrito(carritoDulce)

	assert.NoError(t, err)
	assert.NoError(t, mockDB.ExpectationsWereMet())
}

func TestUpdateDulceInCaritoNotFound(t *testing.T) {
	initialize()

	carritoDulce := getMockCarritoDulce()

	mockDB.ExpectBegin()
	mockDB.ExpectExec(QueryUpdateDulceInCarrito).WillReturnError(gorm.ErrRecordNotFound)
	mockDB.ExpectRollback()

	err := repository.AddDulceInCarrito(carritoDulce)

	typeErr := reflect.TypeOf(err).String()

	assert.Error(t, err)
	assert.Equal(t, "database.NotFoundError", typeErr)
	assert.NoError(t, mockDB.ExpectationsWereMet())
}

func TestUpdateDulceInCaritoInternalServerError(t *testing.T) {
	initialize()
	carritoDulce := getMockCarritoDulce()

	mockDB.ExpectBegin()
	mockDB.ExpectExec(QueryUpdateDulceInCarrito).WillReturnError(gorm.ErrInvalidData)
	mockDB.ExpectRollback()

	err := repository.AddDulceInCarrito(carritoDulce)

	typeErr := reflect.TypeOf(err).String()

	assert.Error(t, err)
	assert.Equal(t, "database.InternalServerError", typeErr)
}

func TestAddDulceInCarritoOK(t *testing.T) {
	initialize()

	carritoDulce := getMockCarritoDulceSinID()

	mockDB.ExpectBegin()
	mockDB.ExpectExec(QueryAddDulceInCarrito).WillReturnResult(sqlmock.NewResult(1, 1))
	mockDB.ExpectCommit()

	err := repository.AddDulceInCarrito(carritoDulce)
	assert.NoError(t, err)
	assert.NoError(t, mockDB.ExpectationsWereMet())
}

func TestAddDulceInCaritoNotFound(t *testing.T) {
	initialize()

	carritoDulce := getMockCarritoDulceSinID()

	mockDB.ExpectBegin()
	mockDB.ExpectExec(QueryAddDulceInCarrito).WillReturnError(gorm.ErrRecordNotFound)
	mockDB.ExpectRollback()

	err := repository.AddDulceInCarrito(carritoDulce)

	typeErr := reflect.TypeOf(err).String()

	assert.Error(t, err)
	assert.Equal(t, "database.NotFoundError", typeErr)
	assert.NoError(t, mockDB.ExpectationsWereMet())
}

func TestAddDulceInCaritoInternalServerError(t *testing.T) {
	initialize()
	carritoDulce := getMockCarritoDulceSinID()

	mockDB.ExpectBegin()
	mockDB.ExpectExec(QueryAddDulceInCarrito).WillReturnError(gorm.ErrInvalidData)
	mockDB.ExpectRollback()

	err := repository.AddDulceInCarrito(carritoDulce)

	typeErr := reflect.TypeOf(err).String()

	assert.Error(t, err)
	assert.Equal(t, "database.InternalServerError", typeErr)
	assert.NoError(t, mockDB.ExpectationsWereMet())
}

func TestDeleteDulceInCarritoOK(t *testing.T) {
	initialize()

	carritoDulce := getMockCarritoDulce()

	mockDB.ExpectBegin()
	mockDB.ExpectExec(QueryDeleteDulceInCarrito).WillReturnResult(sqlmock.NewResult(1, 1))
	mockDB.ExpectCommit()

	err := repository.DeleteDulceInCarrito(carritoDulce)
	assert.NoError(t, err)
	assert.NoError(t, mockDB.ExpectationsWereMet())
}

func TestDeleteDulceInCaritoNotFound(t *testing.T) {
	initialize()

	carritoDulce := getMockCarritoDulce()

	mockDB.ExpectBegin()
	mockDB.ExpectExec(QueryDeleteDulceInCarrito).WillReturnError(gorm.ErrRecordNotFound)
	mockDB.ExpectRollback()

	err := repository.DeleteDulceInCarrito(carritoDulce)

	typeErr := reflect.TypeOf(err).String()

	assert.Error(t, err)
	assert.Equal(t, "database.NotFoundError", typeErr)
	assert.NoError(t, mockDB.ExpectationsWereMet())
}

func TestDeleteDulceInCaritoInternalServerError(t *testing.T) {
	initialize()
	carritoDulce := getMockCarritoDulce()

	mockDB.ExpectBegin()
	mockDB.ExpectExec(QueryDeleteDulceInCarrito).WillReturnError(gorm.ErrInvalidData)
	mockDB.ExpectRollback()

	err := repository.DeleteDulceInCarrito(carritoDulce)

	typeErr := reflect.TypeOf(err).String()

	assert.Error(t, err)
	assert.Equal(t, "database.InternalServerError", typeErr)
	assert.NoError(t, mockDB.ExpectationsWereMet())
}

func TestGetCarritoByCarritoIDErrorNotFound(t *testing.T) {
	initialize()

	mockDB.ExpectQuery(QuerySelectByID).WithArgs(mockCarritoID, 1).WillReturnError(gorm.ErrRecordNotFound)

	carritoRecibido, err := repository.GetByID(mockCarritoID)

	typeErr := reflect.TypeOf(err).String()

	assert.Error(t, err)
	assert.Equal(t, "database.NotFoundError", typeErr)
	assert.Empty(t, carritoRecibido)
}

func TestGetCarritoByCarritoIDInternalServerError(t *testing.T) {
	initialize()

	mockDB.ExpectQuery(QuerySelectByID).WithArgs(mockCarritoID, 1).WillReturnError(gorm.ErrInvalidData)

	carritoRecibido, err := repository.GetByID(mockCarritoID)

	typeErr := reflect.TypeOf(err).String()

	assert.Error(t, err)
	assert.Equal(t, "database.InternalServerError", typeErr)
	assert.Empty(t, carritoRecibido)
}

func TestWhenSaveWasSuccesfullShouldReturnNoError(t *testing.T) {
	initialize()
	mockDB.ExpectBegin()
	mockDB.ExpectExec(queryUpdate).WillReturnResult(sqlmock.NewResult(1, 1))
	mockDB.ExpectCommit()
	mockCarrito := getMockCarrito()

	err := repository.Save(&mockCarrito)

	assert.NoError(t, err)
	assert.NoError(t, mockDB.ExpectationsWereMet())
}

func TestWhenSaveAndCarritoDoesNotContainIDShouldCreateAndReturnNoError(t *testing.T) {
	initialize()
	mockDB.ExpectBegin()
	mockDB.ExpectExec(queryCreate).WillReturnResult(sqlmock.NewResult(1, 1))
	mockDB.ExpectCommit()
	mockCarrito := getMockCarritoToCreate()

	err := repository.Save(&mockCarrito)

	assert.NoError(t, err)
	assert.NoError(t, mockDB.ExpectationsWereMet())
}

func TestWhenSaveWentWrongShouldReturnInternalError(t *testing.T) {
	initialize()
	mockDB.ExpectBegin()
	mockDB.ExpectExec(queryUpdate).WillReturnError(gorm.ErrInvalidData)
	mockDB.ExpectBegin()
	mockCarrito := getMockCarrito()

	err := repository.Save(&mockCarrito)

	assert.Error(t, err)
	assert.Equal(t, "database.InternalServerError", reflect.TypeOf(err).String())
}

func initialize() {
	mockDB, DB = dbmocks.NewDB()
	mockDB.MatchExpectationsInOrder(false)
	repository = Repository{
		DB: DB.Debug(),
	}
}

func getMockCarritoDulce() (carritoDulce entities.CarritoDulce) {
	carritoDulce = entities.CarritoDulce{
		ID:        1,
		CarritoID: 1,
		DulceID:   1,
		Unidades:  2,
		Subtotal:  5900,
	}
	return
}

func getMockCarritoDulceSinID() (carritoDulce entities.CarritoDulce) {
	carritoDulce = entities.CarritoDulce{
		CarritoID: 1,
		DulceID:   1,
		Unidades:  2,
		Subtotal:  5900,
	}
	return
}

func getResponse() (response entities.Carrito) {
	response = entities.Carrito{
		ID:          mockCarritoID,
		PrecioTotal: 1000,
		Subtotal:    995,
		Envio:       5,
		Descuento:   0,
	}
	return
}

func getMockCarrito() entities.Carrito {
	return entities.Carrito{
		ID:          1,
		Subtotal:    1000,
		Descuento:   5,
		Envio:       5,
		PrecioTotal: 1000,
	}
}

func getMockCarritoToCreate() entities.Carrito {
	return entities.Carrito{
		Subtotal:    1000,
		Descuento:   5,
		Envio:       5,
		PrecioTotal: 1000,
	}
}
