package db

import (
	"viandasApp/db"
	"viandasApp/models"
)

func CancelDayOrder(modelDayOrder models.DayOrder, deliveryModel models.Delivery, delete bool) (bool, error) {

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

	if err := tx.Save(&modelDayOrder).Error; err != nil {
		tx.Rollback()
		return false, err
	}

	if delete {

		if err := tx.Exec("DELETE FROM deliveries WHERE id = ?", deliveryModel.ID).Error; err != nil {
			tx.Rollback()
			return false, err
		}

	} else {
		if err := tx.Save(&deliveryModel).Error; err != nil {
			tx.Rollback()
			return false, err
		}
	}

	return true, tx.Commit().Error

}
