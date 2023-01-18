package handlers

import (
	"encoding/json"
	"net/http"

	dbClient "viandasApp/db/client"
	"viandasApp/dtos"

	dbOrder "viandasApp/db/order"
	"viandasApp/handlers"
)

func GetOrders(rw http.ResponseWriter, r *http.Request) {

	usr := handlers.GetUser()

	client, valid := dbClient.CheckExistClient(usr.ID)

	if !valid {
		http.Error(rw, "Error al recuperar el ID del cliente ", http.StatusInternalServerError)
		return
	}

	var responseAllOrdersMenu dtos.OrderViewerResponse

	responseAllOrdersMenu, err := dbOrder.GetAllOrder(client.ID)

	if err != nil {
		http.Error(rw, "Error a recuperar las ordenes de la BD", http.StatusInternalServerError)
		return
	}

	if responseAllOrdersMenu.Order == nil {
		http.Error(rw, "No hay ordenes en la BD", http.StatusNotFound)
	}

	rw.Header().Set("Content-Type", "aplication/json")
	rw.WriteHeader(http.StatusAccepted)
	json.NewEncoder(rw).Encode(responseAllOrdersMenu)

}
