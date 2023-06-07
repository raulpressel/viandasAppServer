package db

import (
	"viandasApp/db"
	"viandasApp/models"
)

/* ChequeoYaExisteUsuario recibe un email de parametro y chequea si ya*/

func CheckExistClient(id string) (models.Client, bool) {
	//ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	//defer cancel()

	db := db.GetDB()

	var client models.Client

	err := db.First(&client, "active = 1 and id_user_kl = ?", id).Error

	if err != nil {
		return client, false
	}

	return client, true
}
