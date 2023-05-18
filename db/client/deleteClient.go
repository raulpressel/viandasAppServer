package db

import (
	"viandasApp/db"
	"viandasApp/models"
)

/* ChequeoYaExisteUsuario recibe un email de parametro y chequea si ya*/

func DeleteClient(clientModel models.Client) (bool, error) {

	db := db.GetDB()

	err := db.Save(&clientModel)

	if err.Error != nil {
		return false, err.Error
	}
	return true, err.Error
}
