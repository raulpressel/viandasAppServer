package db

import (
	"viandasApp/db"
	"viandasApp/models"
)

func GetDayOrdersByOrderId(id int) ([]models.DayOrder, error) {
	db := db.GetDB()

	var dayOrdersModel []models.DayOrder

	err := db.Find(&dayOrdersModel, "order_id = ?", id).Error

	return dayOrdersModel, err

}
