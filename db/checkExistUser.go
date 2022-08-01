package db

import (
	"strconv"

	"viandasApp/models"
	//"go.mongodb.org/mongo-driver/bson"
)

/* ChequeoYaExisteUsuario recibe un email de parametro y chequea si ya*/

func CheckExistUser(email string) (models.User, bool, string) {
	//ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	//defer cancel()

	/* 	db := MysqlCN.Database("twittor")
	   	col := db.Collection("usuarios")

	   	condicion := bson.M{"email": email} */

	//var resultado models.User
	user := models.User{}

	//err := MysqlCN.First(&user, email)
	err := MysqlCN.Where("email = ?", email).First(&user)
	ID := strconv.FormatInt(user.ID, 10)

	if err.Error != nil {

		return user, false, ID

	}
	return user, true, ID
}
