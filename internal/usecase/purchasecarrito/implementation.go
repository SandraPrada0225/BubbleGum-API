package purchasecarrito

import (
	estadoscarrito "bubblegum-api/internal/domain/constants/estados_carrito"
	"bubblegum-api/internal/domain/dto/command/purchasecarrito"
	"bubblegum-api/internal/domain/entities"
	"bubblegum-api/internal/domain/errors/business"
	"bubblegum-api/internal/domain/errors/errormessages"
	"bubblegum-api/internal/repositories/providers"
)

type Implementation struct {
	CarritosProvider providers.CarritosProvider
	VentasProvider   providers.VentasProvider
	UsuariosProvider providers.UsuariosProvider
}

func (useCase Implementation) Execute(command purchasecarrito.PurchaseCarritoCommand) (uint64, error) {
	carrito, err := useCase.CarritosProvider.GetByID(command.CarritoID)
	if err != nil {
		return 0, err
	}

	if carrito.EstadoCarritoID == estadoscarrito.Purchased {
		return 0, business.NewCarritoAlreadyPurchasedError(errormessages.CarritoHasAlreadyBeenPurchased.String())
	}

	usuario, err := useCase.UsuariosProvider.GetByID(command.CompradorID)
	if err != nil {
		return 0, err
	}

	carrito.MarkAsPruchased()

	err = useCase.CarritosProvider.Save(&carrito)
	if err != nil {
		return 0, err
	}

	venta := entities.NewVenta(command.MedioDePagoID, command.CarritoID, command.CompradorID)

	err = useCase.VentasProvider.Create(&venta)
	if err != nil {
		return 0, err
	}

	newCarrito := entities.NewEmptyCarrito()

	err = useCase.CarritosProvider.Save(&newCarrito)
	if err != nil {
		return 0, err
	}

	usuario.UpdateCarritoID(newCarrito.ID)
	err = useCase.UsuariosProvider.Save(&usuario)
	if err != nil {
		return 0, err
	}

	return newCarrito.ID, nil
}
