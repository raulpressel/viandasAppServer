package db

import (
	"viandasApp/db"
	"viandasApp/models"
)

func CancelDayOrder(modelDayOrder models.DayOrder) (bool, error) {

	db := db.GetDB()

	err := db.Save(&modelDayOrder)

	if err.Error != nil {
		return false, err.Error
	}
	return true, err.Error

}
