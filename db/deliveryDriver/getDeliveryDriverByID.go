package db

import (
	"viandasApp/db"
	"viandasApp/models"
)

func GetDeliveryDriverByID(id int) (models.DeliveryDriver, error) {

	db := db.GetDB()

	var deliveryDriverModel models.DeliveryDriver

	err := db.First(&deliveryDriverModel, id).Error

	return deliveryDriverModel, err

}
