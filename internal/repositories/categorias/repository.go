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

const GetDetalleDulcebyCodeSP = "Call GetCategoriasByDulceID(?)"

func (r Repository) GetAll() (categorias []entities.Categoria, err error) {
	err = r.DB.Find(&categorias).Error

	if err != nil {
		params := errormessages.Parameters{
			"resourse": "categorias",
			"error":    err.Error(),
		}
		return []entities.Categoria{}, database.NewInterlServerError(errormessages.InternalServerError.GetMessageWithParams(params))
	}
	return
}

func (r Repository) GetCategoriasByDulceID(dulceID uint64) (categorias []entities.Categoria, err error) {
	err = r.DB.Raw(GetDetalleDulcebyCodeSP, dulceID).Scan(&categorias).Error

	if err != nil {
		params := errormessages.Parameters{
			"resource": "categorias",
		}
		err = database.NewInterlServerError(errormessages.DulceNotFound.GetMessageWithParams(params))
	}
	return
}
