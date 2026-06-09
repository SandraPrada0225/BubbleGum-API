package responses

import "bubblegum-api/internal/domain/entities"

type GetFiltros struct {
	Marcas         []entities.Marca        `json:"marcas"`
	Categorias     []entities.Categoria    `json:"categorias"`
	Presentaciones []entities.Presentacion `json:"presentaciones"`
}
