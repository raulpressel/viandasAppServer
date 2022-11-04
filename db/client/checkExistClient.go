package db

import (
	"viandasApp/db"
	"viandasApp/models"
)

/* ChequeoYaExisteUsuario recibe un email de parametro y chequea si ya*/

func CheckExistClient(id string) (models.Client, bool, string) {
	//ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	//defer cancel()

	/* var db = db.ConnectDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close() */

	db := db.GetDB()

	//user := models.User{}

	var client models.Client

	err := db.First(&client, id).Error

	//err := MysqlCN.First(&user, email)
	//err := db.Where("email = ?", email).First(&user)
	//ID := strconv.FormatInt(user.ID, 10)

	var ID string

	if err != nil {

		return client, false, ID

	}
	return client, true, ID
}
