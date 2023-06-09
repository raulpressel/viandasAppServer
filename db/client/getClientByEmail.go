package db

import (
	"viandasApp/db"
	"viandasApp/models"
)

func GetClientByEmail(email string) (models.Client, error) {

	db := db.GetDB()

	var clientModel models.Client

	db.First(&clientModel, "active = 1 and email = ?", email)

	return clientModel, db.Error

}
