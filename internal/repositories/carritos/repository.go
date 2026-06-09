package carritos

import (
	"bubblegum-api/internal/domain/entities"
	"bubblegum-api/internal/domain/errors/database"
	errormessages "bubblegum-api/internal/domain/errors/error_messages"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func (r Repository) GetByID(carrito_id uint64) (entities.Carrito, error) {
	var carrito entities.Carrito

	err := r.DB.Where("id = ?", carrito_id).Take(&carrito).Error

	if err != nil {
		params := errormessages.Parameters{
			"resource":   "carrito",
			"carrito_id": fmt.Sprint(carrito_id),
		}
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.Carrito{}, database.NewNotFoundError(string(errormessages.CarritoNotFound.GetMessageWithParams(params)))
		}
		params["error"] = err.Error()
		return entities.Carrito{}, database.NewInterlServerError(string(errormessages.InternalServerError.GetMessageWithParams(params)))
	}

	return carrito, nil
}

func (r Repository) GetDulceByCarritoIDAndDulceID(carritoID uint64, dulceID uint64) (carritoDulce entities.CarritoDulce, exists bool, err error) {

	err = r.DB.Where("carrito_id = ? AND dulce_id = ?", carritoID, dulceID).First(&carritoDulce).Error

	if err != nil {
		params := errormessages.Parameters{
			"resource": "carrito",
		}
		if errors.Is(err, gorm.ErrRecordNotFound) {
			exists = false
			err = nil
		} else {
			err = database.NewInterlServerError(errormessages.InternalServerError.GetMessageWithParams(params))
		}
		return
	}
	exists = true
	return
}

func (r Repository) AddDulceInCarrito(carritoDulce entities.CarritoDulce) (err error) {

	fmt.Printf("Guardando: %+v\n", carritoDulce)
	err = r.DB.Save(&carritoDulce).Error
	fmt.Printf("Error Save: %v\n", err)

	if err != nil {
		params := errormessages.Parameters{
			"resource": "carrito",
		}
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = database.NewNotFoundError(errormessages.CarritoDulceNotFound.GetMessageWithParams(params))
		} else {
			err = database.NewInterlServerError(errormessages.InternalServerError.GetMessageWithParams(params))
		}
	}
	return
}

func (r Repository) DeleteDulceInCarrito(carritoDulce entities.CarritoDulce) (err error) {
	err = r.DB.Delete(&carritoDulce).Error

	if err != nil {
		params := errormessages.Parameters{
			"resource": "carrito",
		}
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = database.NewNotFoundError(errormessages.CarritoNotFound.GetMessageWithParams(params))
		} else {
			err = database.NewInterlServerError(errormessages.InternalServerError.GetMessageWithParams(params))
		}
	}
	return
}

func (r Repository) Save(carrito *entities.Carrito) error {
	err := r.DB.Save(carrito).Error
	if err != nil {
		err = database.NewInterlServerError(errormessages.InternalServerError.String())
	}
	return err
}
