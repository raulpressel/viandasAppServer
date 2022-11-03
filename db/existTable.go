package db

import (
	"fmt"
	//"viandasApp/db"
)

/*recibe cualquier MODELO y chequa si existe la tabla en la BD, sino existe crea la tabla en la BD*/

func ExistTable(model interface{}) bool {

	//	var db = ConnectDB()
	/* 	sqlDB, _ := db.DB()
	   	defer sqlDB.Close() */

	db := GetDB()

	if db.Migrator().HasTable(model) {
		fmt.Println("ya existe la tabla", model)

		return false

	}

	db.AutoMigrate(model)

	return true

}
