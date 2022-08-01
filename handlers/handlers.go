package handler

import (
	"log"
	"net/http"
	"os"
	"viandasApp/middlew"
	"viandasApp/routes"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Handlers() {

	router := mux.NewRouter()

	router.HandleFunc("/registro", middlew.CheckDB(routes.Register)).Methods("POST")
	router.HandleFunc("/login", middlew.CheckDB(routes.Login)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
