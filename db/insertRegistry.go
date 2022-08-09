package db

import (
	"viandasApp/models"
)

/*InsertoRegistro es la para final con la BD para insertar los datos del usuario */

func InsertRegistry(userModel models.User) (bool, error) {
	/* ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel() */

	/* db := MysqlCN.Database("twittor")
	col := db.Collection("usuarios") */

	var db = ConnectDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	userModel.Password, _ = EncryptPassword(userModel.Password)
	err := db.Save(&userModel)

	if err.Error != nil {
		return false, err.Error
	}
	return true, err.Error

	/*result, err := col.InsertOne(ctx, u)
	objID, _ := result.InsertedID.(primitive.ObjectID) */

	/* 	return objID.String(), true, nil */

}
