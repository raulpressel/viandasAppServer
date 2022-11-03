package main

import (
	"log"
	db "viandasApp/db"

	h "viandasApp/handlers"
	"viandasApp/routes"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	route, err := h.GetCert("CERT_PEM")
	if err != nil {
		log.Fatal("Key incorrecta")
		return
	}

	key, err := h.GetSecretKey(route)
	if err != nil {
		log.Fatal("Key incorrecta")
		return
	}
	if key == nil {
		log.Fatal("Key incorrecta")
		return
	}

	dsn, err := db.GetKeyDB("DB_CONN")
	if err != nil {
		log.Fatal("Key incorrecta")
		return
	}

	db, err := db.ConnectDB(dsn)

	if err != nil {
		log.Fatal("Sin conexion a la DB")
		return
	}

	if db == nil {
		log.Fatal("Sin conexion a la DB")
		return
	}

	publicDir, err := routes.GetPublicDir("PUBLIC_DIR")
	if err != nil {
		log.Fatal("Key incorrecta")
		return
	}

	routes.Routes(publicDir)

}
