package routes

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"viandasApp/db"
	"viandasApp/handlers"
	address "viandasApp/handlers/address"
	carousel "viandasApp/handlers/carousel"
	categories "viandasApp/handlers/categories"
	city "viandasApp/handlers/cities"
	client "viandasApp/handlers/client"
	deliveryDriver "viandasApp/handlers/deliveryDriver"
	food "viandasApp/handlers/food"
	menu "viandasApp/handlers/menu"
	order "viandasApp/handlers/order"
	pathology "viandasApp/handlers/pathologies"
	tanda "viandasApp/handlers/tanda"

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

	router.HandleFunc("/pathology/getPathology", middlew.CheckDB(middlew.ValidateJWT(pathology.GetAllPathology))).Methods("GET")
	router.HandleFunc("/pathology/uploadPathology", middlew.CheckDB(middlew.ValidateJWTAdmin(pathology.UploadPathology))).Methods("POST")
	router.HandleFunc("/pathology/editPathology", middlew.CheckDB(middlew.ValidateJWTAdmin(pathology.UpdatePathology))).Methods("PUT")
	router.HandleFunc("/pathology/deletePathology", middlew.CheckDB(middlew.ValidateJWTAdmin(pathology.DeletePathology))).Methods("Delete")

	router.HandleFunc("/menu/uploadMenu", middlew.CheckDB(middlew.ValidateJWTAdmin(menu.UploadMenu))).Methods("POST")
	router.HandleFunc("/menu/validateDateMenu", middlew.CheckDB(middlew.ValidateJWTAdmin(menu.ValidateDateMenu))).Methods("POST")
	router.HandleFunc("/menu/editMenu", middlew.CheckDB(middlew.ValidateJWTAdmin(menu.UpdateMenu))).Methods("PUT")
	router.HandleFunc("/menu/deleteMenu", middlew.CheckDB(middlew.ValidateJWTAdmin(menu.DeleteMenu))).Methods("DELETE")
	router.HandleFunc("/menu/getAllMenu", middlew.CheckDB(middlew.ValidateJWTAdmin(menu.GetAllMenu))).Methods("GET")
	router.HandleFunc("/menu/getMenuViewer", middlew.CheckDB(menu.GetMenuViewer)).Methods("GET")
	router.HandleFunc("/menu/getMenuByID", middlew.CheckDB(menu.GetMenuById)).Methods("GET")
	router.HandleFunc("/menu/getMenuByCategory", middlew.CheckDB(menu.GetMenuByCategory)).Methods("GET")
	router.HandleFunc("/menu/getMenuByCategories", middlew.CheckDB(middlew.ValidateJWT(menu.GetMenuByCategories))).Methods("POST")
	router.HandleFunc("/menu/getDayMenu", middlew.CheckDB(menu.GetDayMenuByDate)).Methods("POST")

	router.HandleFunc("/city/getCity", middlew.CheckDB(middlew.ValidateJWT(city.GetAllCities))).Methods("GET")

	router.HandleFunc("/client/registerClient", middlew.CheckDB(middlew.ValidateJWT(client.RegisterClient))).Methods("POST")
	router.HandleFunc("/client/updateClient", middlew.CheckDB(middlew.ValidateJWT(client.UpdateClient))).Methods("POST")
	router.HandleFunc("/client/getClientByIdUser", middlew.CheckDB(middlew.ValidateJWT(client.GetClientByIDUser))).Methods("GET")
	router.HandleFunc("/client/getClient", middlew.CheckDB(middlew.ValidateJWTAdmin(client.GetAllClient))).Methods("GET")
	router.HandleFunc("/client/getClientByTanda", middlew.CheckDB(middlew.ValidateJWTAdmin(client.GetClientsByTandas))).Methods("POST")
	//"/tanda/removeAddressToTanda"

	router.HandleFunc("/address/addAddress", middlew.CheckDB(middlew.ValidateJWT(address.AddAddress))).Methods("POST")
	router.HandleFunc("/address/editAddress", middlew.CheckDB(middlew.ValidateJWT(address.UpdateAddress))).Methods("PUT")
	router.HandleFunc("/address/deleteAddress", middlew.CheckDB(middlew.ValidateJWT(address.DeleteAddress))).Methods("DELETE")
	router.HandleFunc("/address/setFavouriteAddress", middlew.CheckDB(middlew.ValidateJWT(address.SetFavouriteAddress))).Methods("POST")

	router.HandleFunc("/order/uploadOrder", middlew.CheckDB(middlew.ValidateJWT(order.UploadOrder))).Methods("POST")
	router.HandleFunc("/order/getOrderByID", middlew.CheckDB(middlew.ValidateJWT(order.GetOrderById))).Methods("GET")
	router.HandleFunc("/order/getOrderViewer", middlew.CheckDB(middlew.ValidateJWT(order.GetAllOrder))).Methods("GET")
	router.HandleFunc("/order/updateDayOrderAddress", middlew.CheckDB(middlew.ValidateJWT(order.UpdateDayOrderAddress))).Methods("PUT")

	router.HandleFunc("/deliveryDriver/addDeliveryDriver", middlew.CheckDB(middlew.ValidateJWTAdmin(deliveryDriver.UploadDeliveryDriver))).Methods("POST")
	router.HandleFunc("/deliveryDriver/getDeliveryDriver", middlew.CheckDB(middlew.ValidateJWTAdmin(deliveryDriver.GetAllDeliveryDriver))).Methods("GET")
	router.HandleFunc("/deliveryDriver/editDeliveryDriver", middlew.CheckDB(middlew.ValidateJWTAdmin(deliveryDriver.UpdateDeliveryDriver))).Methods("PUT")
	router.HandleFunc("/deliveryDriver/deleteDeliveryDriver", middlew.CheckDB(middlew.ValidateJWTAdmin(deliveryDriver.DeleteDeliveryDriver))).Methods("DELETE")

	router.HandleFunc("/tanda/addTanda", middlew.CheckDB(middlew.ValidateJWTAdmin(tanda.UploadTanda))).Methods("POST")
	router.HandleFunc("/tanda/getTanda", middlew.CheckDB(middlew.ValidateJWTAdmin(tanda.GetAllTanda))).Methods("GET")
	router.HandleFunc("/tanda/editTanda", middlew.CheckDB(middlew.ValidateJWTAdmin(tanda.UpdateTanda))).Methods("PUT")
	router.HandleFunc("/tanda/assignAddressToTanda", middlew.CheckDB(middlew.ValidateJWTAdmin(tanda.AssignAddressToTanda))).Methods("POST")

	var turnMenuModel models.TurnMenu
	var turnModel models.Turn

	var dayModel models.DayMenu

	var pathologyModel models.Pathology

	var tandaAddModel models.TandaAddress

	db.ExistTable(tandaAddModel)

	db.ExistTable(pathologyModel)

	var cityModel models.City

	if db.ExistTable(cityModel) {
		var city = []models.City{{ID: 1, Description: "Paraná", CP: "3100"}, {ID: 2, Description: "San Benito", CP: "3100"}}
		dbc.Create(&city)

	}

	if db.ExistTable(turnModel) {
		var turns = []models.Turn{{ID: 1, Description: "Mediodia"}, {ID: 2, Description: "Noche"}}
		dbc.Create(&turns)

	}
	db.ExistTable(turnMenuModel)
	db.ExistTable(dayModel)

	var addcliModel models.ClientAddress
	var clientModel models.Client
	var addModel models.Address
	var cPModel models.ClientPathology

	db.ExistTable(clientModel)
	db.ExistTable(addModel)
	db.ExistTable(cPModel)
	db.ExistTable(addcliModel)

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
