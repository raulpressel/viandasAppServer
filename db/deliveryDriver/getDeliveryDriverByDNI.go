package db

import (
	"viandasApp/db"
	"viandasApp/models"
)

func GetDeliveryDriverByDNI(dni int) (bool, error) {

	db := db.GetDB()

	var deliveryDriverModel models.DeliveryDriver

	err := db.First(&deliveryDriverModel, "dni = ?", dni).Error

	if err != nil {
		return false, err
	}

	return true, err

}
