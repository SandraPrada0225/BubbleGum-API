package puchasecarrito

import (
	command "bubblegum-api/internal/domain/dto/command/purchasecarrito"
	contract "bubblegum-api/internal/domain/dto/contracts/purchasecarrito"
	"bubblegum-api/internal/domain/dto/responses"
	"bubblegum-api/internal/domain/errors"
	errormessages "bubblegum-api/internal/domain/errors/errormessages"
	"bubblegum-api/internal/domain/errors/rest"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PurchaseCarritoUseCase interface {
	Execute(command command.PurchaseCarritoCommand) (uint64, error)
}

type PurchaseCarrito struct {
	UseCase PurchaseCarritoUseCase
}

func (handler PurchaseCarrito) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request contract.Request
		id, err := strconv.ParseUint(c.Param("id"), 10, 64)

		if err != nil {
			params := errormessages.Parameters{
				"error": err.Error(),
			}
			c.JSON(http.StatusBadRequest, rest.NewBadRequestError(errormessages.InvalidTypeerror.GetMessageWithParams(params)).Error())
			return
		}

		request.URLParams.CarritoID = id
		err = c.ShouldBindJSON(&request.Body)

		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		command := command.NewPurchaseCarritoCommandFromRequest(request)
		fmt.Printf("AFTER BIND: %+v\n", command)
		fmt.Printf("Command: %+v\n", command)
		body, _ := io.ReadAll(c.Request.Body)
		fmt.Println("este es el body ", string(body))
		newCarritoID, err := handler.UseCase.Execute(command)
		if err != nil {
			errors.GetAPIErrors(err, c)
			return
		}

		c.JSON(http.StatusOK, responses.NewResponse(newCarritoID))
	}
}
