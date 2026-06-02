package providers

import (
	"bubblegum-api/internal/domain/dto/query"
)

type DulcesProvider interface {
	GetByCode(codigo string) (dulce query.DetalleDulce, err error)
}
