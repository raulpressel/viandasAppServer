package main

/* "github.com/gorilla/mux" */

import (
	"log"
	db "viandasApp/db"
	handler "viandasApp/handlers"
)

func main() {

	if db.CheckConnection() == 0 {
		log.Fatal("Sin conexion a la DB")
		return
	} else {

		handler.Handlers()

	}

}
