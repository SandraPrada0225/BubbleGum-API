package carritos

import (
	"bubblegum-api/internal/domain/entities"
	"bubblegum-api/internal/domain/errors/database"
	errormessages "bubblegum-api/internal/domain/errors/error_messages"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func (r Repository) GetCarritoByCarritoID(carrito_id uint64) (entities.Carrito, error) {
	var carrito entities.Carrito

	err := r.DB.Where("id = ?", carrito_id).Take(&carrito).Error

	if err != nil {
		params := errormessages.Parameters{
			"resource":   "carrito",
			"carrito_id": fmt.Sprint(carrito_id),
		}
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.Carrito{}, database.NewNotFoundError(string(errormessages.CarritoNotFound.GetMessageWithParams(params)))
		}
		params["error"] = err.Error()
		return entities.Carrito{}, database.NewInterlServerError(string(errormessages.InternalServerError.GetMessageWithParams(params)))
	}

	return carrito, nil
}
