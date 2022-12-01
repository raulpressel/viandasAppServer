package db

import (
	"viandasApp/db"
	"viandasApp/models"
)

func UploadOrder(orderModel models.Order, dayOrderModel []models.DayOrder) (bool, error) {

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

	if err := tx.Save(&orderModel).Error; err != nil {
		tx.Rollback()
		return false, err
	}

	for i := range dayOrderModel {

		dayOrderModel[i].OrderID = orderModel.ID
	}

	if err := tx.CreateInBatches(&dayOrderModel, len(dayOrderModel)).Error; err != nil {
		tx.Rollback()
		return false, err
	}

	return true, tx.Commit().Error

}
