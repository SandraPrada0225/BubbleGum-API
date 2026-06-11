package errormessages

import "fmt"

type (
	ErrorMessage string
	Parameters   map[string]interface{}
)

const (
	CarritoDulceNotFound           ErrorMessage = "No se encontró un detalle carrito_dulce con ese codigo"
	DulceNotFound                  ErrorMessage = "No se encontró un dulce con ese codigo"
	UsuarioNotFound                ErrorMessage = "No se encontró un usuario con este id"
	CarritoNotBelonging            ErrorMessage = "El carrito pertenece a otro usuario"
	InternalServerError            ErrorMessage = "Ha ocurrido un error inesperado"
	IdMustBeAPositiveNumber        ErrorMessage = "El ID deber ser un número positivo"
	UnitLimitExceded               ErrorMessage = "Las unidades requeridas exceden las disponibles"
	CarritoNotFound                ErrorMessage = "No se encontró un carrito con ese id"
	InvalidTypeerror               ErrorMessage = "El tipo de dato es invalido"
	CarritoHasAlreadyBeenPurchased ErrorMessage = "El carrito ya ha sido comprado"
)

func (e ErrorMessage) String() string {
	return string(e)
}

func (e ErrorMessage) GetMessageWithParams(params Parameters) string {
	msg := e.String()

	for key, value := range params {
		msg = fmt.Sprintf("%s. %s: %s", msg, key, value)
	}
	return msg
}
