package getfiltros

import (
	"bubblegum-api/internal/domain/dto/query"
	"bubblegum-api/internal/repositories/providers"
)

type Implementation struct {
	CategoriasProvider     providers.CategoriasProvider
	MarcasProvider         providers.MarcasProvider
	PresentacionesProvider providers.PresentacionesProvider
}

func (UseCase Implementation) Execute() (query.GetFiltros, error) {

	marcas, err := UseCase.MarcasProvider.GetAll()

	if err != nil {
		return query.GetFiltros{}, err
	}

	categorias, err := UseCase.CategoriasProvider.GetAll()

	if err != nil {
		return query.GetFiltros{}, err
	}

	presentaciones, err := UseCase.PresentacionesProvider.GetAll()

	if err != nil {
		return query.GetFiltros{}, err
	}

	return query.GetFiltros{
		Marcas:         marcas,
		Categorias:     categorias,
		Presentaciones: presentaciones,
	}, nil
}
