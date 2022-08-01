package db

import (
	"viandasApp/models"
)

/*InsertoRegistro es la para final con la BD para insertar los datos del usuario */

func InsertRegistry(u models.User) (string, bool, error) {
	/* ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel() */

	/* db := MysqlCN.Database("twittor")
	col := db.Collection("usuarios") */

	u.Password, _ = EncriptarPassword(u.Password)
	MysqlCN.Save(&u)

	return "", true, nil

	/* 	result, err := col.InsertOne(ctx, u)
	   	if err != nil {
	   		return "", false, err
	   	}

	   	objID, _ := result.InsertedID.(primitive.ObjectID) */

	/* 	return objID.String(), true, nil */

}
