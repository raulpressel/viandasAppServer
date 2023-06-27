package db

import (
	"viandasApp/db"
	"viandasApp/models"
)

func CheckEmptyDeliveryDrivers() ([]models.Delivery, error) {

	db := db.GetDB()

	var deliveries []models.Delivery

	err := db.Find(&deliveries, "delivery_driver_id is null and address_id <> 100").Error

	if err != nil {
		return deliveries, err
	}

	return deliveries, nil
}
