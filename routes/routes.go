package routes

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"viandasApp/db"
	"viandasApp/handlers"
	carousel "viandasApp/handlers/carousel"
	categories "viandasApp/handlers/categories"
	food "viandasApp/handlers/food"
	login "viandasApp/handlers/login"
	menu "viandasApp/handlers/menu"
	register "viandasApp/handlers/register"

	"viandasApp/middlew"
	"viandasApp/models"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Routes() {

	router := mux.NewRouter()

	router.HandleFunc("/isAuthorizated", handlers.IsAuthorizated).Methods("GET")

	router.HandleFunc("/register", middlew.CheckDB(register.Register)).Methods("POST")
	router.HandleFunc("/login", middlew.CheckDB(login.Login)).Methods("POST")

	router.HandleFunc("/carrousel/uploadBanner", middlew.CheckDB(middlew.ValidateJWT(carousel.UploadBanner))).Methods("POST")
	router.HandleFunc("/carrousel/editBanner", middlew.CheckDB(middlew.ValidateJWT(carousel.UpdateBanner))).Methods("PUT")
	router.HandleFunc("/carrousel/deleteBanner", middlew.CheckDB(middlew.ValidateJWT(carousel.DeleteBanner))).Methods("DELETE")

	router.HandleFunc("/carrousel/getBannersIndex", middlew.CheckDB(carousel.GetBanners)).Methods("GET")
	router.HandleFunc("/carrousel/getBanners", middlew.CheckDB(middlew.ValidateJWT(carousel.GetAllBanners))).Methods("GET")

	router.HandleFunc("/food/uploadFood", middlew.CheckDB(middlew.ValidateJWT(food.UploadFood))).Methods("POST")
	router.HandleFunc("/food/editFood", middlew.CheckDB(middlew.ValidateJWT(food.UpdateFood))).Methods("PUT")
	router.HandleFunc("/food/deleteFood", middlew.CheckDB(middlew.ValidateJWT(food.DeleteFood))).Methods("DELETE")
	router.HandleFunc("/food/getFood", middlew.CheckDB(middlew.ValidateJWT(food.GetAllFood))).Methods("GET")
	router.HandleFunc("/food/getFoodByCategory", middlew.CheckDB(middlew.ValidateJWT(food.GetFoodByCategory))).Methods("GET")

	router.HandleFunc("/category/getCategory", middlew.CheckDB(middlew.ValidateJWT(categories.GetAllCategories))).Methods("GET")
	router.HandleFunc("/category/uploadCategory", middlew.CheckDB(middlew.ValidateJWT(categories.UploadCategory))).Methods("POST")
	router.HandleFunc("/category/editCategory", middlew.CheckDB(middlew.ValidateJWT(categories.UpdateCategory))).Methods("PUT")
	router.HandleFunc("/category/deleteCategory", middlew.CheckDB(middlew.ValidateJWT(categories.DeelteCategory))).Methods("Delete")

	router.HandleFunc("/menu/uploadMenu", middlew.CheckDB(middlew.ValidateJWT(menu.UploadMenu))).Methods("POST")
	router.HandleFunc("/menu/editMenu", middlew.CheckDB(middlew.ValidateJWT(menu.UpdateMenu))).Methods("PUT")
	router.HandleFunc("/menu/deleteMenu", middlew.CheckDB(middlew.ValidateJWT(menu.DeleteMenu))).Methods("DELETE")
	router.HandleFunc("/menu/getMenu", middlew.CheckDB(menu.GetMenu)).Methods("GET")
	router.HandleFunc("/menu/getMenuByCategory", middlew.CheckDB(menu.GetMenuByCategory)).Methods("GET")

	var turnMenu models.TurnMenu

	var dayModel models.DayMenu

	db.ExistTable(turnMenu)
	db.ExistTable(dayModel)

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	fmt.Println(PORT)

	router.PathPrefix("/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("/var/www/default/htdocs/public/"))))

	handler := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
