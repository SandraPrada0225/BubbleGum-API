package dulces

import (
	"bubblegum-api/internal/domain/dto/query"
	"bubblegum-api/internal/domain/errors/database"
	errormessages "bubblegum-api/internal/domain/errors/error_messages"
	"errors"

	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

const GetDetalleDulcebyCodeSP = "Call GetDetalleDulceByCode(?)"

func (r Repository) GetByCode(codigo string) (detalleDulce query.DetalleDulce, err error) {
	err = r.DB.Raw(GetDetalleDulcebyCodeSP, codigo).Take(&detalleDulce).Error

	if err != nil {
		params := errormessages.Parameters{
			"resource":   "dulces",
			"dulce_code": codigo,
		}
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = database.NewNotFoundError(errormessages.DulceNotFound.GetMessageWithParams(params))
		} else {
			err = database.NewInterlServerError(errormessages.DulceNotFound.GetMessageWithParams(params))
		}
	}
	return
}
