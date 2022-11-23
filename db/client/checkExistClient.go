package db

import (
	"viandasApp/db"
	"viandasApp/models"
)

/* ChequeoYaExisteUsuario recibe un email de parametro y chequea si ya*/

func CheckExistClient(id string) (models.Client, error) {
	//ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	//defer cancel()

	db := db.GetDB()

	var client models.Client

	//err := db.First(&client, id).Error

	err := db.First(&client, "id_user_kl = ?", id).Error

	return client, err
}
