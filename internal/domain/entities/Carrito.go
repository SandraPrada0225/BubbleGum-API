package entities

import estadocarrito "bubblegum-api/internal/domain/constants/estados_carrito"

type Carrito struct {
	ID              uint64
	PrecioTotal     float64
	Subtotal        float64
	Envio           float64
	Descuento       float64
	EstadoCarritoID uint64
}

func (carrito *Carrito) MarkAsPruchased() {
	carrito.EstadoCarritoID = estadocarrito.Purchased
}

func NewEmptyCarrito() Carrito {
	return Carrito{
		EstadoCarritoID: estadocarrito.Active,
	}
}
