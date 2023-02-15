package db

import (
	"viandasApp/db"
	"viandasApp/dtos"
	"viandasApp/models"
)

func UploadOrder(orderModel models.Order, dayOrderModel []models.DayOrder) (bool, error, *dtos.SaveOrderResponse) {

	var response dtos.SaveOrderResponse

	db := db.GetDB()

	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return false, err, nil
	}

	if err := tx.Save(&orderModel).Error; err != nil {
		tx.Rollback()
		return false, err, nil
	}

	for i := range dayOrderModel {

		dayOrderModel[i].OrderID = orderModel.ID
	}

	if err := tx.CreateInBatches(&dayOrderModel, len(dayOrderModel)).Error; err != nil {
		tx.Rollback()
		return false, err, nil
	}

	response.IDOrder = orderModel.ID

	return true, tx.Commit().Error, &response

}
