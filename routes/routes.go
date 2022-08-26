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

	router.HandleFunc("/isAuthorizated", handlers.IsAuthorizated).Methods("GET")

	router.HandleFunc("/register", middlew.CheckDB(handlers.Register)).Methods("POST")
	router.HandleFunc("/login", middlew.CheckDB(handlers.Login)).Methods("POST")

	router.HandleFunc("/uploadBanner", middlew.CheckDB(middlew.ValidateJWT(handlers.UploadBanner))).Methods("POST")

	router.HandleFunc("/carrousel/getBannersIndex", middlew.CheckDB(handlers.GetBanners)).Methods("GET")
	router.HandleFunc("/carrousel/getBanners", middlew.CheckDB(middlew.ValidateJWT(handlers.GetAllBanners))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "3000"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
