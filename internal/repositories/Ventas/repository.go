package ventas

import (
	"bubblegum-api/internal/domain/dto/responses"
	"bubblegum-api/internal/domain/entities"
	"bubblegum-api/internal/domain/errors/database"
	"bubblegum-api/internal/domain/errors/errormessages"
	"errors"

	"gorm.io/gorm"
)

const GetListByUserIDSP = "Call GetPurchaseListByUserID(?)"

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

func (r Repository) GetListByUserID(userID uint64) (responses.GetPurchaseList, error) {
	var purchaseList responses.GetPurchaseList

	err := r.DB.Raw(GetListByUserIDSP, userID).Find(&purchaseList.PurchaseList).Error
	if err != nil {
		params := errormessages.Parameters{
			"resource": "ventas",
			"user_id":  userID,
			"error":    err.Error(),
		}
		return responses.GetPurchaseList{}, database.NewInterlServerError(errormessages.InternalServerError.GetMessageWithParams(params))
	}
	return purchaseList, nil
}
