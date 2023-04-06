package db

import (
	"viandasApp/db"
	"viandasApp/models"
)

func GetStatusOrder(id int) (models.StatusOrder, error) {
	db := db.GetDB()

	var modelStatusOrder models.StatusOrder

	err := db.First(&modelStatusOrder, id).
		Where("id <> 0").Error

	return modelStatusOrder, err

}
