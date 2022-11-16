package db

import (
	"viandasApp/db"
	"viandasApp/models"
)

func UploadPathology(pathologyModel models.Pathology) (bool, error) {

	db := db.GetDB()

	err := db.Save(&pathologyModel)

	if err.Error != nil {
		return false, err.Error
	}
	return true, err.Error

}
