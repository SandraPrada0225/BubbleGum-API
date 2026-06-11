package purchasecarrito

type Request struct {
	Body      Body
	URLParams URLParams
}

type Body struct {
	CompradorID   uint64 `json:"comprador_id" validate:"required,number"`
	MedioDePagoID uint64 `json:"medio_de_pago_id" validate:"required,number"`
}

type URLParams struct {
	CarritoID uint64
}
