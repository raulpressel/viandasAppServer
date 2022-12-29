package db

import (
	"viandasApp/db"
	"viandasApp/models"
)

func UpdateDayOrderAddress(dayOrderModel models.DayOrder) (bool, error) {

	db := db.GetDB()

	err := db.Save(&dayOrderModel)

	if err.Error != nil {
		return false, err.Error
	}
	return true, err.Error

}
