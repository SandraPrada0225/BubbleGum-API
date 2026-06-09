package dulces

import (
	"bubblegum-api/internal/domain/dto/responses"
	"bubblegum-api/internal/domain/entities"
	"bubblegum-api/internal/domain/errors/database"
	errormessages "bubblegum-api/internal/domain/errors/error_messages"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

const (
	GetDetalleDulceByIDSP   = "Call GetDetalleDulceByID(?)"
	GetDetalleDulcebyCodeSP = "Call GetDetalleDulceByCode(?)"
)

type Repository struct {
	DB *gorm.DB
}

func (r Repository) GetByCode(codigo string) (detalleDulce responses.DetalleDulce, err error) {
	err = r.DB.Raw(GetDetalleDulcebyCodeSP, codigo).Take(&detalleDulce).Error

	if err != nil {
		params := errormessages.Parameters{
			"resource":   "dulces",
			"dulce_code": codigo,
		}
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = database.NewNotFoundError(errormessages.DulceNotFound.GetMessageWithParams(params))
		} else {
			err = database.NewInterlServerError(errormessages.InternalServerError.GetMessageWithParams(params))
		}
	}
	return
}

func (r Repository) GetByID(id uint64) (dulce entities.Dulce, err error) {
	err = r.DB.Where("id = ?", id).Take(&dulce).Error

	if err != nil {
		params := errormessages.Parameters{
			"resource": "dulce",
		}
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = database.NewNotFoundError(errormessages.DulceNotFound.GetMessageWithParams(params))
		} else {
			err = database.NewInterlServerError(errormessages.InternalServerError.GetMessageWithParams(params))
		}
	}
	return
}

func (r Repository) GetDetailByID(id uint64) (detalleDulce responses.DetalleDulce, err error) {
	err = r.DB.Raw(GetDetalleDulceByIDSP, id).Take(&detalleDulce).Error

	if err != nil {
		params := errormessages.Parameters{
			"resource": "dulces",
			"dulce_id": fmt.Sprint(id),
		}
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = database.NewNotFoundError(errormessages.DulceNotFound.GetMessageWithParams(params))
		} else {
			err = database.NewInterlServerError(errormessages.InternalServerError.GetMessageWithParams(params))
		}
	}
	return
}

func (r Repository) GetDulcesListByCarritoID(carrito_id uint64) ([]entities.CarritoDulce, error) {
	var dulcesInCarrito []entities.CarritoDulce

	err := r.DB.Model(&entities.CarritoDulce{}).Where("carrito_id=?", carrito_id).Find(&dulcesInCarrito).Error

	if err != nil {
		params := errormessages.Parameters{
			"resource":   "dulces",
			"carrito_id": carrito_id,
			"error":      err.Error(),
		}
		return []entities.CarritoDulce{}, database.NewInterlServerError(errormessages.InternalServerError.GetMessageWithParams(params))
	}
	return dulcesInCarrito, nil
}
