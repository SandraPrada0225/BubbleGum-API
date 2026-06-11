package ventas

import (
	"bubblegum-api/internal/domain/entities"
	"bubblegum-api/internal/domain/errors/database"
	"bubblegum-api/internal/domain/errors/errormessages"
	"errors"

	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func (r Repository) Create(venta *entities.Venta) error {

	err := r.DB.Create(venta).Error

	if err != nil {
		params := errormessages.Parameters{
			"resource":   "ventas",
			"carrito_id": venta.CarritoID,
		}

		if errors.Is(gorm.ErrForeignKeyViolated, err) {
			return database.NewNotFoundError(errormessages.CarritoNotFound.GetMessageWithParams(params))
		}

		if errors.Is(gorm.ErrDuplicatedKey, err) {
			return database.NewConflictError(errormessages.CarritoHasAlreadyBeenPurchased.GetMessageWithParams(params))
		}
		params["error"] = err.Error()

		return database.NewInterlServerError(errormessages.InternalServerError.GetMessageWithParams(params))
	}
	return nil
}
