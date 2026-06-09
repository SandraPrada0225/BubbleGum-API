package handlers

import (
	"bubblegum-api/internal/domain/dto/responses"
	"bubblegum-api/internal/domain/errors/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetDulcebyCode struct {
	UseCase UseCase
}

type UseCase interface {
	Execute(codigo string) (responses.DetalleDulce, error)
}

func (handler GetDulcebyCode) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		codigo := c.Param("codigo")
		dulce, err := handler.UseCase.Execute(codigo)
		if err != nil {
			switch err.(type) {
			case database.NotFoundError:
				c.JSON(http.StatusNotFound, err.Error())
			default:
				c.JSON(http.StatusInternalServerError, err.Error())
			}
			return
		}
		c.JSON(http.StatusOK, dulce)
	}
}
