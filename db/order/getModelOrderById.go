package db

import (
	"viandasApp/db"
	"viandasApp/models"
)

func GetModelOrderById(id int) (models.Order, error) {
	db := db.GetDB()

	var orderModel models.Order

	err := db.First(&orderModel, id).Error

	return orderModel, err

}
