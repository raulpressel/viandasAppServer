package db

import (
	"viandasApp/db"
	"viandasApp/models"
)

func UpdateDeliveryDriver(deliveryDriverModel models.DeliveryDriver, vehicleModel models.Vehicle, addressModel models.Address) (bool, error) {

	db := db.GetDB()

	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return false, err
	}

	if err := tx.Save(&addressModel).Error; err != nil {
		tx.Rollback()
		return false, err
	}

	if err := tx.Save(&vehicleModel).Error; err != nil {
		tx.Rollback()
		return false, err
	}

	if err := tx.Save(&deliveryDriverModel).Error; err != nil {
		tx.Rollback()
		return false, err
	}

	return true, tx.Commit().Error
}
