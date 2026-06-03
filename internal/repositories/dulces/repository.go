package dulces

import (
	"bubblegum-api/internal/domain/dto/query"
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

func (r Repository) GetByCode(codigo string) (detalleDulce query.DetalleDulce, err error) {
	err = r.DB.Raw(GetDetalleDulcebyCodeSP, codigo).Take(&detalleDulce).Error

	if err != nil {
		params := errormessages.Parameters{
			"resource":   "dulces",
			"dulce_code": codigo,
		}
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = database.NewNotFoundError(errormessages.DulceNotFound.GetMessageWithParams(params))
		} else {
			err = database.NewInterlServerError(errormessages.DulceNotFound.GetMessageWithParams(params))
		}
	}
	return
}

func (r Repository) GetDetailByID(id uint64) (detalleDulce query.DetalleDulce, err error) {
	err = r.DB.Raw(GetDetalleDulceByIDSP, id).Take(&detalleDulce).Error

	if err != nil {
		params := errormessages.Parameters{
			"resource":   "dulces",
			"dulce_code": fmt.Sprint(id),
		}
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = database.NewNotFoundError(errormessages.DulceNotFound.GetMessageWithParams(params))
		} else {
			err = database.NewInterlServerError(errormessages.InternalServerError.GetMessageWithParams(params))
		}
	}
	return
}

func (r Repository) GetDulcesListByCarritoID(carrito_id uint64) ([]uint64, error) {
	var dulcesIDList []uint64

	err := r.DB.Model(&entities.CarritoDulce{}).Select("dulces_id").
		Where("carritos_id = ?", carrito_id).Find(&dulcesIDList).Error

	if err != nil {
		params := errormessages.Parameters{
			"resource":   "carrito",
			"carrito_id": carrito_id,
			"error":      err.Error(),
		}
		return []uint64{}, database.NewInterlServerError(string(errormessages.InternalServerError.GetMessageWithParams(params)))
	}
	return dulcesIDList, nil
}
