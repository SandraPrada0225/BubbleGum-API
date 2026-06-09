package entities

type Carrito struct {
	ID          uint64
	PrecioTotal float64
	Subtotal    float64
	Envio       float64
	Descuento   float64
}
