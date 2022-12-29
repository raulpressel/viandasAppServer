package db

import (
	"viandasApp/db"
	"viandasApp/models"
)

func GetDayOrderById(id int) (models.DayOrder, error) {
	db := db.GetDB()

	var dayOrderModel models.DayOrder

	err := db.First(&dayOrderModel, id).Error

	return dayOrderModel, err

}
