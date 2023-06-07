package db

import (
	"viandasApp/db"
	"viandasApp/models"
)

func GetDayOrdersByOrderId(order_id int) ([]models.DayOrder, error) {
	db := db.GetDB()

	var dayOrdersModel []models.DayOrder

	err := db.Where("order_id = ? AND status = 1", order_id).Find(&dayOrdersModel).Error

	return dayOrdersModel, err

}
