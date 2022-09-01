package db

import (
	"strconv"

	"viandasApp/db"
	"viandasApp/models"
)

/* ChequeoYaExisteUsuario recibe un email de parametro y chequea si ya*/

func CheckExistUser(email string) (models.User, bool, string) {
	//ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	//defer cancel()

	var db = db.ConnectDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	user := models.User{}

	//err := MysqlCN.First(&user, email)
	err := db.Where("email = ?", email).First(&user)
	ID := strconv.FormatInt(user.ID, 10)

	if err.Error != nil {

		return user, false, ID

	}
	return user, true, ID
}
