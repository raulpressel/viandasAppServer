package db

import (
	"viandasApp/db"
	"viandasApp/models"
)

func UploadZone(zoneModel models.Zone) (bool, error) {

	db := db.GetDB()

	err := db.Save(&zoneModel)

	if err.Error != nil {
		return false, err.Error
	}
	return true, err.Error

}
