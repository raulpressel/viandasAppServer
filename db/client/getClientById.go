package db

import (
	"viandasApp/db"
	"viandasApp/models"
)

func GetClientById(id int) (models.Client, error) {

	db := db.GetDB()

	var clientModel models.Client

	err := db.First(&clientModel, id).Error

	return clientModel, err

}
