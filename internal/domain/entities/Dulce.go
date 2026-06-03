package entities

import "time"

type Dulce struct {
	ID               uint64    `json: "id"`
	Nombre           string    `json: "nombre"`
	MarcaID          uint64    `json: "marca_id"`
	Precio           float64   `json: "precio"`
	Peso             float64   `json: "peso"`
	Unidades         int       `json: "unidades"`
	PresentacionID   uint64    `json: "presentacion_id"`
	Descripcion      string    `json: "descripcion"`
	Imagen           string    `json: "imagen"`
	FechaVencimiento time.Time `json: "fecha_vencimiento"`
	FechaExpedicion  time.Time `json: "fecha_expedicion"`
	Disponibles      int       `json: "disponibles"`
	Codigo           string    `json: "codigo"`
}
