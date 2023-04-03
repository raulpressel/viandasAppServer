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

	/* clientModel, err := dbClient.GetClientById(orderModel.ClientID)

	if err != nil {
		return false, err, nil
	}

	clientModel.Observation = orderModel.Observation

	if err := tx.Save(&clientModel).Error; err != nil {
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

	response.IDOrder = orderModel.ID */

	return true, tx.Commit().Error

}
