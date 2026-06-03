package providers

import (
	"bubblegum-api/internal/domain/dto/query"
)

type DulcesProvider interface {
	GetByCode(codigo string) (dulce query.DetalleDulce, err error)
	GetDetailByID(id uint64) (dulce query.DetalleDulce, err error)
	GetDulcesListByCarritoID(carrito_id uint64) ([]uint64, error)
}
