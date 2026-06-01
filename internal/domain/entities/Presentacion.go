package entities

type Presentacion struct {
	ID     int    `json: "id"`
	Nombre string `json: "nombre"`
}

func (Presentacion) TableName() string {
	return "Presentaciones"
}
