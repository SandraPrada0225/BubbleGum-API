package database

type InternalServerError struct {
	mensaje string
}

func (e InternalServerError) Error() string {
	return e.mensaje
}

func NewInterlServerError(mensaje string) error {
	return InternalServerError{
		mensaje: mensaje,
	}
}
