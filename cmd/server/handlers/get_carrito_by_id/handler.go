package getcarritobyid

import (
	"bubblegum-api/internal/domain/dto/query"
	"bubblegum-api/internal/domain/errors/database"
	errormessages "bubblegum-api/internal/domain/errors/error_messages"
	"bubblegum-api/internal/domain/errors/rest"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetCarritoByID struct {
	UseCase UseCase
}

type UseCase interface {
	Execute(id uint64) (query.GetDetalleCarrito, error)
}

func (handler GetCarritoByID) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			params := errormessages.Parameters{
				"error": err.Error(),
			}
			c.JSON(http.StatusBadRequest, rest.NewBadRequestError(errormessages.InvalidTypeerror.GetMessageWithParams(params)).Error())
			return
		}

		carrito, err := handler.UseCase.Execute(id)
		if err != nil {
			switch err.(type) {
			case database.NotFoundError:
				c.JSON(http.StatusNotFound, err.Error())
			default:
				c.JSON(http.StatusInternalServerError, err.Error())
			}
			return
		}
		c.JSON(http.StatusOK, carrito)
	}
}
