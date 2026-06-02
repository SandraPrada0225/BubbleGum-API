package providers

import (
	"bubblegum-api/internal/domain/entities"
)

type CategoriasProvider interface {
	GetCategoriasByDulceID(dulceID uint64) (categorias []entities.Categoria, err error)
	GetAll() (categorias []entities.Categoria, err error)
}
