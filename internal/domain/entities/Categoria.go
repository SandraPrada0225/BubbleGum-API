package entities

type Categoria struct {
	ID     int    `json: "id"`
	Nombre string `json: "nombre"`
}

func (Categoria) TableName() string {
	return "Categorias"
}
