package db

import (
	"viandasApp/db"
	"viandasApp/models"
)

func GetZoneById(id int) (models.Zone, error) {

	db := db.GetDB()

	var zoneModel models.Zone

	err := db.First(&zoneModel, id).Error

	return zoneModel, err

}
