package entities

type Marca struct {
	ID     int    `json: "id"`
	Nombre string `json: "nombre"`
}

func (Marca) TableName() string {
	return "Marcas"
}
