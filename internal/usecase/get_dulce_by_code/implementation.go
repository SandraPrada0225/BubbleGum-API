package getdulcebycode

import (
	"bubblegum-api/internal/domain/entities"
	"bubblegum-api/internal/repositories/providers"
)

type Implementation struct {
	DulcesProvider providers.DulcesProvider
}

func (UseCase Implementation) Execute(codigo string) (entities.Dulce, error) {
	return UseCase.DulcesProvider.GetByCode(codigo)
}
