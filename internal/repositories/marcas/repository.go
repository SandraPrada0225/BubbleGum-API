package marcas

import (
	"bubblegum-api/internal/domain/entities"
	"bubblegum-api/internal/domain/errors/database"
	errormessages "bubblegum-api/internal/domain/errors/error_messages"

	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func (r Repository) GetAll() (marcas []entities.Marca, err error) {
	err = r.DB.Find(&marcas).Error

	if err != nil {
		params := errormessages.Parameters{
			"resource": "marcas",
			"error":    err.Error(),
		}

		return []entities.Marca{},
			database.NewInterlServerError(errormessages.InternalServerError.GetMessageWithParams(params))
	}
	return
}
