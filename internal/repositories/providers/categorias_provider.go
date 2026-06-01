package providers

import "bubblegum-api/internal/domain/entities"

type CategoriasProvider interface {
	GetAll() (categorias []entities.Categoria, err error)
}
