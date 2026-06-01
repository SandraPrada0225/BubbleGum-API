package providers

import "bubblegum-api/internal/domain/entities"

type PresentacionesProvider interface {
	GetAll() (presentaciones []entities.Presentacion, err error)
}
