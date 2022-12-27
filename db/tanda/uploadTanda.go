package db

import (
	"viandasApp/db"
	"viandasApp/models"
)

func UploadTanda(tandaModel models.Tanda) (bool, error) {

	db := db.GetDB()

	err := db.Save(&tandaModel)

	if err.Error != nil {
		return false, err.Error
	}
	return true, err.Error

}
