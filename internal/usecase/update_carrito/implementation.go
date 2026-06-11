package updatecarrito

import (
	"bubblegum-api/internal/domain/constants"

	"bubblegum-api/internal/domain/dto/contracts/updatecarrito"
	"bubblegum-api/internal/domain/dto/responses"
	"bubblegum-api/internal/domain/entities"
	"bubblegum-api/internal/domain/errors/business"
	"bubblegum-api/internal/domain/errors/errormessages"

	"bubblegum-api/internal/repositories/providers"
	"fmt"
)

type Implementation struct {
	CarritoProvider providers.CarritosProvider
	DulcesProvider  providers.DulcesProvider
}

func (useCase Implementation) Execute(carritoID uint64, movements updatecarrito.Body) (responses.MovementsResult, error) {
	_, err := useCase.CarritoProvider.GetByID(carritoID)
	if err != nil {
		return responses.MovementsResult{}, err
	}

	movementsResult := responses.NewMovementsResult()

	for index, movememnt := range movements.Movements {
		var operationResult constants.CarritoOperationResult
		carritoDulce, exists, err := useCase.CarritoProvider.GetDulceByCarritoIDAndDulceID(carritoID, movememnt.DulceID)
		fmt.Printf("exists=%v carritoDulce=%+v\n", exists, carritoDulce)

		if err != nil {
			movementsResult.AddResult(index, movememnt.DulceID, constants.Error.String(), err.Error())
			continue
		}
		switch {
		case movememnt.Unidades == 0:
			operationResult = constants.Deleted
			err = useCase.CarritoProvider.DeleteDulceInCarrito(carritoDulce)
			if err != nil {
				movementsResult.AddResult(index, movememnt.DulceID, constants.Error.String(), err.Error())
				continue
			}
		case !exists:
			operationResult = constants.Created
			carritoDulce = entities.NewCarritoDulce(carritoID, movememnt.DulceID)
			err := useCase.save(movememnt, carritoDulce)
			if err != nil {
				movementsResult.AddResult(index, movememnt.DulceID, constants.Error.String(), err.Error())
				continue
			}
		case exists:
			operationResult = constants.Updated
			err := useCase.save(movememnt, carritoDulce)
			if err != nil {
				movementsResult.AddResult(index, movememnt.DulceID, constants.Error.String(), err.Error())
				continue
			}
		}
		movementsResult.AddResult(index, movememnt.DulceID, operationResult.String(), "")
	}
	return movementsResult, err
}

func (UseCase Implementation) save(movement updatecarrito.Movement, carritoDulce entities.CarritoDulce) error {
	dulce, err := UseCase.DulcesProvider.GetByID(movement.DulceID)
	if err != nil {
		return err
	}
	if movement.Unidades > dulce.Disponibles {
		return business.NewUnitLimitExceded(errormessages.UnitLimitExceded.String())
	}
	carritoDulce = entities.UpdateCarritoDulce(carritoDulce, movement.Unidades, dulce.Precio)

	err = UseCase.CarritoProvider.AddDulceInCarrito(carritoDulce)
	return err
}
