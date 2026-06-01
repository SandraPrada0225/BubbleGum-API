package presentaciones

import (
	"bubblegum-api/internal/domain/entities"
	"bubblegum-api/internal/domain/errors/database"
	errormessages "bubblegum-api/internal/domain/errors/error_messages"

	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func (r Repository) GetAll() (presentaciones []entities.Presentacion, err error) {

	err = r.DB.Find(&presentaciones).Error

	if err != nil {
		return []entities.Presentacion{}, database.NewInterlServerError(errormessages.InternalServerError)
	}
	return
}
