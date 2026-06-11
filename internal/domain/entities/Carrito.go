package entities

import estadoscarrito "bubblegum-api/internal/domain/constants/estados_carrito"

type Carrito struct {
	ID              uint64
	PrecioTotal     float64
	Subtotal        float64
	Envio           float64
	Descuento       float64
	EstadoCarritoID uint64 `gorm:"column:estados_carrito_id"`
}

func (carrito *Carrito) MarkAsPruchased() {
	carrito.EstadoCarritoID = estadoscarrito.Purchased
}

func NewEmptyCarrito() Carrito {
	return Carrito{
		EstadoCarritoID: estadoscarrito.Active,
	}
}
