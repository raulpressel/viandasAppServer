package db

import (
	"viandasApp/db"
	"viandasApp/models"
)

func GetVehicleByID(id int) (models.Vehicle, error) {

	db := db.GetDB()

	var vehicleModel models.Vehicle

	err := db.First(&vehicleModel, id).Error

	return vehicleModel, err

}
