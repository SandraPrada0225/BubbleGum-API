package providers

import (
	"bubblegum-api/internal/domain/dto/responses"
	"bubblegum-api/internal/domain/entities"
)

type DulcesProvider interface {
	GetByCode(codigo string) (dulce responses.DetalleDulce, err error)
	GetByID(id uint64) (dulce entities.Dulce, err error)
	GetDetailByID(id uint64) (dulce responses.DetalleDulce, err error)
	GetDulcesListByCarritoID(carrito_id uint64) ([]entities.CarritoDulce, error)
}
