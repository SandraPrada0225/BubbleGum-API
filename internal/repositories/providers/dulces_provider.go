package providers

import "bubblegum-api/internal/domain/entities"

type DulcesProvider interface {
	GetByCode(codigo string) (dulce entities.Dulce, err error)
}
