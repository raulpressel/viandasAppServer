package db

import "fmt"

/*recibe cualquier MODELO y chequa si existe la tabla en la BD, sino existe crea la tabla en la BD*/

func ExistTable(model interface{}) {

	var db = ConnectDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	//user := userDto.ToModelUser()

	if db.Migrator().HasTable(model) {
		fmt.Println("ya existe la tabla", model)

	} else {
		fmt.Println("creamos la tabla", model)
		db.AutoMigrate(model)
	}

}
