package ventas

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

const (
	createQuery                 = "INSERT INTO `ventas` (`medio_de_pago_id`,`carrito_id`,`comprador_id`,`created_at`) VALUES (?,?,?,?)"
	queryGetListByUserID        = "Call GetPurchaseListByUserID(?)"
	mockUserID           uint64 = 4213
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

func TestGetPurchaseListWhenIsSuccesfullShouldReturnList(t *testing.T) {
	initialize()

	puchaseList := getMockedPurchaseList()

	mockDB.ExpectQuery(queryGetListByUserID).WithArgs(mockUserID).WillReturnRows(
		sqlmock.NewRows([]string{
			"id",
			"fecha",
			"medio_de_pago_id",
			"medio_de_pago",
			"carrito_id",
			"precio_total",
			"subtotal",
			"descuento",
			"envio",
			"estado_carrito_id",
			"estado_carrito",
		}).AddRow(
			puchaseList.PurchaseList[0].ID,
			puchaseList.PurchaseList[0].Fecha,
			puchaseList.PurchaseList[0].MedioDePagoID,
			puchaseList.PurchaseList[0].MedioDePago,
			puchaseList.PurchaseList[0].CarritoID,
			puchaseList.PurchaseList[0].PrecioTotal,
			puchaseList.PurchaseList[0].Subtotal,
			puchaseList.PurchaseList[0].Descuento,
			puchaseList.PurchaseList[0].Envio,
			puchaseList.PurchaseList[0].EstadoCarritoID,
			puchaseList.PurchaseList[0].EstadoCarrito).
			AddRow(
				puchaseList.PurchaseList[1].ID,
				puchaseList.PurchaseList[1].Fecha,
				puchaseList.PurchaseList[1].MedioDePagoID,
				puchaseList.PurchaseList[1].MedioDePago,
				puchaseList.PurchaseList[1].CarritoID,
				puchaseList.PurchaseList[1].PrecioTotal,
				puchaseList.PurchaseList[1].Subtotal,
				puchaseList.PurchaseList[1].Descuento,
				puchaseList.PurchaseList[1].Envio,
				puchaseList.PurchaseList[1].EstadoCarritoID,
				puchaseList.PurchaseList[1].EstadoCarrito),
	)
	response, err := repository.GetListByUserID(mockUserID)

	assert.NoError(t, err)
	assert.Equal(t, puchaseList, response)
	assert.NoError(t, mockDB.ExpectationsWereMet())
}

func TestGetPurchaseListWhenWentWrongShouldReturnInternalServerError(t *testing.T) {
	initialize()

	mockDB.ExpectQuery(queryGetListByUserID).WithArgs(mockUserID).WillReturnError(gorm.ErrInvalidData)

	response, err := repository.GetListByUserID(mockUserID)
	typeErr := reflect.TypeOf(err).String()

	assert.Error(t, err)
	assert.Equal(t, "database.InternalServerError", typeErr)
	assert.Empty(t, response)
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

func getMockedPurchaseList() responses.GetPurchaseList {
	fecha := time.Date(2023, 10, 17, 0, 0, 0, 0, time.Local)
	return responses.GetPurchaseList{
		PurchaseList: []responses.Purchase{
			{
				ID:              1,
				Fecha:           fecha,
				MedioDePagoID:   1,
				MedioDePago:     "contraentrega",
				PrecioTotal:     100,
				Subtotal:        97,
				Descuento:       2,
				Envio:           5,
				EstadoCarritoID: 1,
				EstadoCarrito:   "comprador",
			},
			{
				ID:              2,
				Fecha:           fecha,
				MedioDePagoID:   1,
				MedioDePago:     "credito",
				PrecioTotal:     120,
				Subtotal:        100,
				Descuento:       0,
				Envio:           20,
				EstadoCarritoID: 1,
				EstadoCarrito:   "comprador",
			},
		},
	}
}
