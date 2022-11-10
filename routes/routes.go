package routes

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"viandasApp/db"
	"viandasApp/handlers"
	carousel "viandasApp/handlers/carousel"
	categories "viandasApp/handlers/categories"
	food "viandasApp/handlers/food"
	menu "viandasApp/handlers/menu"
	pathology "viandasApp/handlers/pathologies"

	"viandasApp/middlew"
	"viandasApp/models"

	"github.com/gorilla/mux"

	"github.com/rs/cors"
)

func Routes(publicDir string) {

	dbc := db.GetDB()
	if dbc == nil {
		fmt.Println("conexion perdida")
		return
	}
	router := mux.NewRouter()

	router.HandleFunc("/isAuthorizated", handlers.IsAuthorizated).Methods("GET")

	//router.HandleFunc("/administration", middlew.CheckDB(middlew.ValidateJWTAdmin(carousel.GetAllBanners))).Methods("GET")

	router.HandleFunc("/carrousel/uploadBanner", middlew.CheckDB(middlew.ValidateJWTAdmin(carousel.UploadBanner))).Methods("POST")
	router.HandleFunc("/carrousel/editBanner", middlew.CheckDB(middlew.ValidateJWTAdmin(carousel.UpdateBanner))).Methods("PUT")
	router.HandleFunc("/carrousel/deleteBanner", middlew.CheckDB(middlew.ValidateJWTAdmin(carousel.DeleteBanner))).Methods("DELETE")
	router.HandleFunc("/carrousel/getBannersIndex", middlew.CheckDB(carousel.GetBanners)).Methods("GET")
	router.HandleFunc("/carrousel/getBanners", middlew.CheckDB(middlew.ValidateJWTAdmin(carousel.GetAllBanners))).Methods("GET")

	router.HandleFunc("/food/uploadFood", middlew.CheckDB(middlew.ValidateJWTAdmin(food.UploadFood))).Methods("POST")
	router.HandleFunc("/food/editFood", middlew.CheckDB(middlew.ValidateJWTAdmin(food.UpdateFood))).Methods("PUT")
	router.HandleFunc("/food/deleteFood", middlew.CheckDB(middlew.ValidateJWTAdmin(food.DeleteFood))).Methods("DELETE")
	router.HandleFunc("/food/getFood", middlew.CheckDB(middlew.ValidateJWTAdmin(food.GetAllFood))).Methods("GET")
	router.HandleFunc("/food/getFoodByCategory", middlew.CheckDB(middlew.ValidateJWTAdmin(food.GetFoodByCategory))).Methods("GET")
	router.HandleFunc("/food/getImageByCategory", middlew.CheckDB(food.GetImageByCategory)).Methods("GET")

	router.HandleFunc("/category/getCategory", middlew.CheckDB(categories.GetAllCategories)).Methods("GET")
	router.HandleFunc("/category/uploadCategory", middlew.CheckDB(middlew.ValidateJWTAdmin(categories.UploadCategory))).Methods("POST")
	router.HandleFunc("/category/editCategory", middlew.CheckDB(middlew.ValidateJWTAdmin(categories.UpdateCategory))).Methods("PUT")
	router.HandleFunc("/category/deleteCategory", middlew.CheckDB(middlew.ValidateJWTAdmin(categories.DeleteCategory))).Methods("Delete")

	router.HandleFunc("/pathology/getPathology", middlew.CheckDB(middlew.ValidateJWTAdmin(pathology.GetAllPathology))).Methods("GET")
	router.HandleFunc("/pathology/uploadCategory", middlew.CheckDB(middlew.ValidateJWTAdmin(pathology.UploadPathology))).Methods("POST")
	router.HandleFunc("/pathology/editCategory", middlew.CheckDB(middlew.ValidateJWTAdmin(pathology.UpdatePathology))).Methods("PUT")
	router.HandleFunc("/pathology/deleteCategory", middlew.CheckDB(middlew.ValidateJWTAdmin(pathology.DeletePathology))).Methods("Delete")

	router.HandleFunc("/menu/uploadMenu", middlew.CheckDB(middlew.ValidateJWTAdmin(menu.UploadMenu))).Methods("POST")
	router.HandleFunc("/menu/validateDateMenu", middlew.CheckDB(middlew.ValidateJWTAdmin(menu.ValidateDateMenu))).Methods("POST")
	router.HandleFunc("/menu/editMenu", middlew.CheckDB(middlew.ValidateJWTAdmin(menu.UpdateMenu))).Methods("PUT")
	router.HandleFunc("/menu/deleteMenu", middlew.CheckDB(middlew.ValidateJWTAdmin(menu.DeleteMenu))).Methods("DELETE")
	router.HandleFunc("/menu/getAllMenu", middlew.CheckDB(middlew.ValidateJWTAdmin(menu.GetAllMenu))).Methods("GET")
	router.HandleFunc("/menu/getMenu", middlew.CheckDB(menu.GetMenu)).Methods("GET")
	router.HandleFunc("/menu/getMenuByID", middlew.CheckDB(menu.GetMenuById)).Methods("GET")
	router.HandleFunc("/menu/getMenuByCategory", middlew.CheckDB(menu.GetMenuByCategory)).Methods("GET")
	router.HandleFunc("/menu/getDayMenu", middlew.CheckDB(menu.GetDayMenuByDate)).Methods("POST")

	var turnMenuModel models.TurnMenu
	var turnModel models.Turn

	var dayModel models.DayMenu

	var pathologyModel models.Pathology

	if db.ExistTable(pathologyModel) {
		var patho = []models.Pathology{{ID: 1, Description: "Diabetes", Active: true}, {ID: 2, Description: "Hipertension", Active: true}}
		dbc.Create(&patho)

	}

	if db.ExistTable(turnModel) {
		var turns = []models.Turn{{ID: 1, Description: "Mediodia"}, {ID: 2, Description: "Noche"}}
		dbc.Create(&turns)

	}
	db.ExistTable(turnMenuModel)
	db.ExistTable(dayModel)

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	router.PathPrefix("/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir(publicDir))))

	handler := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}

func GetPublicDir(key string) (string, error) {
	publicDir := os.Getenv(key)
	if publicDir == "" {
		return publicDir, errors.New("missing key")
	}
	return publicDir, nil
}
