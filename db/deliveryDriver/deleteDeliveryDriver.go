package db

import (
	"viandasApp/db"
	"viandasApp/models"
)

func DeleteDeliveryDriver(deliveryDriverModel models.DeliveryDriver) (bool, error) {

	db := db.GetDB()

	err := db.Save(&deliveryDriverModel).Error

	if err != nil {
		return false, err
	}
	return true, err

}
