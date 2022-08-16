package routes

import (
	"log"
	"net/http"
	"os"
	"viandasApp/handlers"
	"viandasApp/middlew"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Routes() {

	router := mux.NewRouter()

	router.HandleFunc("/registro", middlew.CheckDB(handlers.Register)).Methods("POST")
	router.HandleFunc("/login", middlew.CheckDB(handlers.Login)).Methods("POST")

	router.HandleFunc("/uploadBanner", middlew.CheckDB(middlew.ValidateJWT(handlers.UploadBanner))).Methods("POST")
	//	router.HandleFunc("/getBanners", middlew.CheckDB(middlew.ValidateJWT(handlers.UploadBanner))).Methods("POST")
	//router.HandleFunc("/obtenerAvatar", middlew.ChequeoBD(middlew.ValidoJWT(routes.ObtenerAvatar))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
