package providers

import (
	"bubblegum-api/internal/domain/dto/query"
	"bubblegum-api/internal/domain/entities"
)

type DulcesProvider interface {
	GetByCode(codigo string) (dulce query.DetalleDulce, err error)
	GetDetailByID(id uint64) (dulce query.DetalleDulce, err error)
	GetDulcesListByCarritoID(carrito_id uint64) ([]entities.CarritoDulce, error)
}
