package providers

import "bubblegum-api/internal/domain/entities"

type VentasProvider interface {
	Create(venta *entities.Venta) error
}
