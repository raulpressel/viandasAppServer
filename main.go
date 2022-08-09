package main

/* "github.com/gorilla/mux" */

import (
	"log"
	db "viandasApp/db"
	"viandasApp/routes"
)

func main() {

	if db.CheckConnection() == 0 {
		log.Fatal("Sin conexion a la DB")
		return
	} else {

		/* _db := db.ConnectDB()
		_db.AutoMigrate(banner) */

		routes.Routes()
		/* 		var model models.User
		   		fmt.Println(getType(model))

		   		Teststruct(model)

		   		var model2 models.User2

		   		Teststruct(model2) */
	}

}

/* func getType(myvar models.User) string {
	return reflect.TypeOf(myvar).String()
}

func Teststruct(x interface{}) {
	// type switch
	switch x.(type) {
	case models.User:
		fmt.Println("User")
	case models.User2:
		fmt.Println("int type")
	default:
		fmt.Println("Error")
	}
}
*/
