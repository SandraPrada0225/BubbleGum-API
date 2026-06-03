package entities

type Carrito struct {
	ID          uint64  `json: "id"`
	PrecioTotal float64 `json: "precio_total"`
	Subtotal    float64 `json: "subtotal"`
	Envio       float64 `json: "envio"`
	Descuento   float64 `json: "descuento"`
}
