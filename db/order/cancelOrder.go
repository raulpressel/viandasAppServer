package db

import (
	"viandasApp/db"
	"viandasApp/models"
)

func CancelOrder(modelOrder models.Order) (bool, error) {

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

	if err := tx.Save(&modelOrder).Error; err != nil {
		tx.Rollback()
		return false, err
	}

	modelDayOrder, err := GetDayOrdersByOrderIdandDate(modelOrder.ID)

	if err != nil {
		return false, err
	}

	for i := range modelDayOrder {

		modelDayOrder[i].Status = false

	}

	if err := tx.Save(&modelDayOrder).Error; err != nil {
		tx.Rollback()
		return false, err
	}

	if err := tx.Exec("DELETE FROM deliveries WHERE order_id = ?", modelOrder.ID).Error; err != nil {
		tx.Rollback()
		return false, err
	}

	return true, tx.Commit().Error

}
