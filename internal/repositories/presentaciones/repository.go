package presentaciones

import (
	"bubblegum-api/internal/domain/entities"
	"bubblegum-api/internal/domain/errors/database"
	"bubblegum-api/internal/domain/errors/errormessages"

	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func (r Repository) GetAll() (presentaciones []entities.Presentacion, err error) {

	err = r.DB.Find(&presentaciones).Error

	if err != nil {
		params := errormessages.Parameters{
			"resource": "presentaciones",
			"error":    err.Error(),
		}

		return []entities.Presentacion{},
			database.NewInterlServerError(errormessages.InternalServerError.GetMessageWithParams(params))
	}
	return
}
