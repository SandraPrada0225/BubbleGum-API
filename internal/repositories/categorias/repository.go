package categorias

import (
	"bubblegum-api/internal/domain/entities"
	"bubblegum-api/internal/domain/errors/database"
	errormessages "bubblegum-api/internal/domain/errors/error_messages"

	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func (r Repository) GetAll() (categorias []entities.Categoria, err error) {
	err = r.DB.Find(&categorias).Error

	if err != nil {
		return []entities.Categoria{}, database.NewInterlServerError(errormessages.InternalServerError)
	}
	return
}
