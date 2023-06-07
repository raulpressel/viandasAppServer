package db

import (
	"viandasApp/db"
	"viandasApp/models"
)

func UpdateOrder(modelOrder models.Order) (bool, error) {

	db := db.GetDB()

	err := db.Save(&modelOrder)

	if err.Error != nil {
		return false, err.Error
	}
	return true, err.Error

}
