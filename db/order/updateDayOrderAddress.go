package db

import (
	"viandasApp/db"
	"viandasApp/models"
)

func UpdateDayOrderAddress(dayOrdersModel []models.DayOrder, deliveryModel *models.Delivery, orderModel models.Order) (bool, error) {

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

	for _, dayOrder := range dayOrdersModel {
		if err := tx.Model(&models.DayOrder{}).Where("id = ?", dayOrder.ID).Update("address_id", dayOrder.AddressID).Error; err != nil {
			tx.Rollback()
			return false, err
		}
	}

	if deliveryModel != nil {

		if err := tx.Save(deliveryModel).Error; err != nil {
			tx.Rollback()
			return false, err
		}

		if err := tx.Model(&models.Order{}).Where("id = ?", orderModel.ID).Update("total", orderModel.Total).Error; err != nil {
			tx.Rollback()
			return false, err
		}
	}

	return true, tx.Commit().Error

}
