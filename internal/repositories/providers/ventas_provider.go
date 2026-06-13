package providers

import (
	"bubblegum-api/internal/domain/dto/responses"
	"bubblegum-api/internal/domain/entities"
)

type VentasProvider interface {
	Create(venta *entities.Venta) error
	GetListByUserID(userID uint64) (responses.GetPurchaseList, error)
}
