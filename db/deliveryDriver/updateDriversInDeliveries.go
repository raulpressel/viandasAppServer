package db

import (
	"viandasApp/db"
	"viandasApp/models"
)

func UpdateDriversInDeliveries(deliveries []models.Delivery) (bool, error) {

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

	if err := tx.Save(&deliveries).Error; err != nil {
		tx.Rollback()
		return false, err
	}

	return true, tx.Commit().Error

}
