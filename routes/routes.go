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
	setting "viandasApp/handlers/setting"
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

	router.HandleFunc("/app/isAuthorizated", handlers.IsAuthorizated).Methods("GET")

	//router.HandleFunc("/app/administration", middlew.CheckDB(middlew.ValidateJWTAdmin(carousel.GetAllBanners))).Methods("GET")

	router.HandleFunc("/app/carrousel/uploadBanner", middlew.CheckDB(middlew.ValidateJWTAdmin(carousel.UploadBanner))).Methods("POST")
	router.HandleFunc("/app/carrousel/editBanner", middlew.CheckDB(middlew.ValidateJWTAdmin(carousel.UpdateBanner))).Methods("PUT")
	router.HandleFunc("/app/carrousel/deleteBanner", middlew.CheckDB(middlew.ValidateJWTAdmin(carousel.DeleteBanner))).Methods("DELETE")
	router.HandleFunc("/app/carrousel/getBannersIndex", middlew.CheckDB(carousel.GetBanners)).Methods("GET")
	router.HandleFunc("/app/carrousel/getBanners", middlew.CheckDB(middlew.ValidateJWTAdmin(carousel.GetAllBanners))).Methods("GET")

	router.HandleFunc("/app/food/uploadFood", middlew.CheckDB(middlew.ValidateJWTAdmin(food.UploadFood))).Methods("POST")
	router.HandleFunc("/app/food/editFood", middlew.CheckDB(middlew.ValidateJWTAdmin(food.UpdateFood))).Methods("PUT")
	router.HandleFunc("/app/food/deleteFood", middlew.CheckDB(middlew.ValidateJWTAdmin(food.DeleteFood))).Methods("DELETE")
	router.HandleFunc("/app/food/getFood", middlew.CheckDB(middlew.ValidateJWTAdmin(food.GetAllFood))).Methods("GET")
	router.HandleFunc("/app/food/getFoodByCategory", middlew.CheckDB(middlew.ValidateJWTAdmin(food.GetFoodByCategory))).Methods("GET")
	router.HandleFunc("/app/food/getImageByCategory", middlew.CheckDB(food.GetImageByCategory)).Methods("GET")

	router.HandleFunc("/app/category/getCategory", middlew.CheckDB(categories.GetAllCategories)).Methods("GET")
	router.HandleFunc("/app/category/uploadCategory", middlew.CheckDB(middlew.ValidateJWTAdmin(categories.UploadCategory))).Methods("POST")
	router.HandleFunc("/app/category/editCategory", middlew.CheckDB(middlew.ValidateJWTAdmin(categories.UpdateCategory))).Methods("PUT")
	router.HandleFunc("/app/category/deleteCategory", middlew.CheckDB(middlew.ValidateJWTAdmin(categories.DeleteCategory))).Methods("Delete")

	router.HandleFunc("/app/pathology/getPathology", middlew.CheckDB(middlew.ValidateJWT(pathology.GetAllPathology))).Methods("GET")
	router.HandleFunc("/app/pathology/uploadPathology", middlew.CheckDB(middlew.ValidateJWTAdmin(pathology.UploadPathology))).Methods("POST")
	router.HandleFunc("/app/pathology/editPathology", middlew.CheckDB(middlew.ValidateJWTAdmin(pathology.UpdatePathology))).Methods("PUT")
	router.HandleFunc("/app/pathology/deletePathology", middlew.CheckDB(middlew.ValidateJWTAdmin(pathology.DeletePathology))).Methods("Delete")

	router.HandleFunc("/app/menu/uploadMenu", middlew.CheckDB(middlew.ValidateJWTAdmin(menu.UploadMenu))).Methods("POST")
	router.HandleFunc("/app/menu/validateDateMenu", middlew.CheckDB(middlew.ValidateJWTAdmin(menu.ValidateDateMenu))).Methods("POST")
	router.HandleFunc("/app/menu/editMenu", middlew.CheckDB(middlew.ValidateJWTAdmin(menu.UpdateMenu))).Methods("PUT")
	router.HandleFunc("/app/menu/deleteMenu", middlew.CheckDB(middlew.ValidateJWTAdmin(menu.DeleteMenu))).Methods("DELETE")
	router.HandleFunc("/app/menu/getAllMenu", middlew.CheckDB(middlew.ValidateJWTAdmin(menu.GetAllMenu))).Methods("GET")
	router.HandleFunc("/app/menu/getMenuViewer", middlew.CheckDB(menu.GetMenuViewer)).Methods("POST")
	router.HandleFunc("/app/menu/getMenuByID", middlew.CheckDB(menu.GetMenuById)).Methods("GET")
	router.HandleFunc("/app/menu/getMenuByCategory", middlew.CheckDB(menu.GetMenuByCategory)).Methods("GET")
	router.HandleFunc("/app/menu/getMenuByCategories", middlew.CheckDB(middlew.ValidateJWT(menu.GetMenuByCategories))).Methods("POST")
	router.HandleFunc("/app/menu/getDayMenu", middlew.CheckDB(menu.GetDayMenuByDate)).Methods("POST")

	router.HandleFunc("/app/city/getCity", middlew.CheckDB(middlew.ValidateJWT(city.GetAllCities))).Methods("GET")

	router.HandleFunc("/app/client/registerClient", middlew.CheckDB(middlew.ValidateJWT(client.RegisterClient))).Methods("POST")
	router.HandleFunc("/app/client/updateClient", middlew.CheckDB(middlew.ValidateJWT(client.UpdateClient))).Methods("POST")
	router.HandleFunc("/app/client/getClientByIdUser", middlew.CheckDB(middlew.ValidateJWT(client.GetClientByIDUser))).Methods("GET")
	router.HandleFunc("/app/client/getClient", middlew.CheckDB(middlew.ValidateJWTAdmin(client.GetAllClient))).Methods("GET")
	router.HandleFunc("/app/client/getClientByTanda", middlew.CheckDB(middlew.ValidateJWTAdmin(client.GetClientsByTandas))).Methods("POST")
	router.HandleFunc("/app/client/addNote", middlew.CheckDB(middlew.ValidateJWTAdmin(client.AddClientNote))).Methods("POST")
	router.HandleFunc("/app/client/editNote", middlew.CheckDB(middlew.ValidateJWTAdmin(client.EditClientNote))).Methods("POST")

	router.HandleFunc("/app/address/addAddress", middlew.CheckDB(middlew.ValidateJWT(address.AddAddress))).Methods("POST")
	router.HandleFunc("/app/address/editAddress", middlew.CheckDB(middlew.ValidateJWT(address.UpdateAddress))).Methods("PUT")
	router.HandleFunc("/app/address/deleteAddress", middlew.CheckDB(middlew.ValidateJWT(address.DeleteAddress))).Methods("DELETE")
	router.HandleFunc("/app/address/setFavouriteAddress", middlew.CheckDB(middlew.ValidateJWT(address.SetFavouriteAddress))).Methods("POST")

	router.HandleFunc("/app/order/uploadOrder", middlew.CheckDB(middlew.ValidateJWT(order.UploadOrder))).Methods("POST")
	router.HandleFunc("/app/order/getOrderByID", middlew.CheckDB(middlew.ValidateJWT(order.GetOrderById))).Methods("GET")
	router.HandleFunc("/app/order/getOrderViewer", middlew.CheckDB(middlew.ValidateJWT(order.GetAllOrder))).Methods("GET")
	router.HandleFunc("/app/order/getOrderByIdClient", middlew.CheckDB(middlew.ValidateJWTAdmin(order.GetOrderByIdClient))).Methods("GET")
	router.HandleFunc("/app/order/updateDayOrderAddress", middlew.CheckDB(middlew.ValidateJWT(order.UpdateDayOrderAddress))).Methods("GET")
	router.HandleFunc("/app/order/getOrders", middlew.CheckDB(middlew.ValidateJWTAdmin(order.GetOrders))).Methods("POST")

	router.HandleFunc("/app/deliveryDriver/addDeliveryDriver", middlew.CheckDB(middlew.ValidateJWTAdmin(deliveryDriver.UploadDeliveryDriver))).Methods("POST")
	router.HandleFunc("/app/deliveryDriver/getDeliveryDriver", middlew.CheckDB(middlew.ValidateJWTAdmin(deliveryDriver.GetAllDeliveryDriver))).Methods("GET")
	router.HandleFunc("/app/deliveryDriver/editDeliveryDriver", middlew.CheckDB(middlew.ValidateJWTAdmin(deliveryDriver.UpdateDeliveryDriver))).Methods("PUT")
	router.HandleFunc("/app/deliveryDriver/deleteDeliveryDriver", middlew.CheckDB(middlew.ValidateJWTAdmin(deliveryDriver.DeleteDeliveryDriver))).Methods("DELETE")

	router.HandleFunc("/app/tanda/addTanda", middlew.CheckDB(middlew.ValidateJWTAdmin(tanda.UploadTanda))).Methods("POST")
	router.HandleFunc("/app/tanda/getTanda", middlew.CheckDB(middlew.ValidateJWTAdmin(tanda.GetAllTanda))).Methods("GET")
	router.HandleFunc("/app/tanda/editTanda", middlew.CheckDB(middlew.ValidateJWTAdmin(tanda.UpdateTanda))).Methods("PUT")
	router.HandleFunc("/app/tanda/assignAddressToTanda", middlew.CheckDB(middlew.ValidateJWTAdmin(tanda.AssignAddressToTanda))).Methods("POST")
	router.HandleFunc("/app/tanda/removeAddressToTanda", middlew.CheckDB(middlew.ValidateJWTAdmin(tanda.RemoveAddressToTanda))).Methods("POST")
	router.HandleFunc("/app/tanda/deleteTanda", middlew.CheckDB(middlew.ValidateJWTAdmin(tanda.DeleteTanda))).Methods("DELETE")

	router.HandleFunc("/app/setting/addDiscount", middlew.CheckDB(middlew.ValidateJWTAdmin(setting.UploadDiscount))).Methods("POST")
	router.HandleFunc("/app/setting/getDiscount", middlew.CheckDB(middlew.ValidateJWTAdmin(setting.GetAllDiscount))).Methods("GET")
	router.HandleFunc("/app/setting/editDiscount", middlew.CheckDB(middlew.ValidateJWTAdmin(setting.UpdateDiscount))).Methods("PUT")
	router.HandleFunc("/app/setting/deleteDiscount", middlew.CheckDB(middlew.ValidateJWTAdmin(setting.DeleteDiscount))).Methods("DELETE")

	router.HandleFunc("/app/setting/addZone", middlew.CheckDB(middlew.ValidateJWTAdmin(setting.UploadZone))).Methods("POST")
	router.HandleFunc("/app/setting/getZone", middlew.CheckDB(middlew.ValidateJWTAdmin(setting.GetAllZone))).Methods("GET")
	router.HandleFunc("/app/setting/editZone", middlew.CheckDB(middlew.ValidateJWTAdmin(setting.UpdateZone))).Methods("PUT")
	router.HandleFunc("/app/setting/deleteZone", middlew.CheckDB(middlew.ValidateJWTAdmin(setting.DeleteZone))).Methods("DELETE")

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
