package providers

import "bubblegum-api/internal/domain/entities"

type MarcasProvider interface {
	GetAll() (marcas []entities.Marca, err error)
}
