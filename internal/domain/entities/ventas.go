package entities

import "time"

type Venta struct {
	ID            uint64
	MedioDePagoID uint64 `gorm:"column:medios_de_pago_id"`
	CarritoID     uint64 `gorm:"column:carritos_id"`
	CompradorID   uint64 `gorm:"column:usuarios_id"`
	CreatedAt     time.Time
}

func (Venta) TableName() string {
	return "ventas"
}

func NewVenta(medioDePago, carritoID, compradorID uint64) Venta {
	return Venta{
		MedioDePagoID: medioDePago,
		CarritoID:     carritoID,
		CompradorID:   compradorID,
	}
}
