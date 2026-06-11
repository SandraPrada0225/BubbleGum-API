package usuarios

import (
	"errors"

	"bubblegum-api/internal/domain/entities"
	"bubblegum-api/internal/domain/errors/database"
	"bubblegum-api/internal/domain/errors/errormessages"

	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func (r Repository) Save(usuario *entities.Usuario) error {
	err := r.DB.Save(&usuario).Error

	if err != nil {
		params := errormessages.Parameters{
			"resource": "usuarios",
		}

		if errors.Is(gorm.ErrDuplicatedKey, err) {

			return database.NewConflictError(errormessages.CarritoNotBelonging.String())
		}
		return database.NewInterlServerError(errormessages.InternalServerError.GetMessageWithParams(params))
	}
	return nil
}

func (r Repository) GetByID(usuarioID uint64) (entities.Usuario, error) {
	var usuario entities.Usuario
	err := r.DB.Where("id = ?", usuarioID).Take(&usuario).Error

	if err != nil {
		params := errormessages.Parameters{
			"resource":   "usuarios",
			"usuario_id": usuarioID,
		}

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.Usuario{}, database.NewNotFoundError(errormessages.UsuarioNotFound.GetMessageWithParams(params))
		}
		params["error"] = err.Error()

		return entities.Usuario{}, database.NewInterlServerError(errormessages.InternalServerError.GetMessageWithParams(params))
	}
	return usuario, nil
}
