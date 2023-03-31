package handlers

import (
	"encoding/json"
	"net/http"

	dbClient "viandasApp/db/client"

	dbOrder "viandasApp/db/order"
	"viandasApp/handlers"
)

func GetOrderViewer(rw http.ResponseWriter, r *http.Request) {

	usr := handlers.GetUser()

	client, valid := dbClient.CheckExistClient(usr.ID)

	if !valid {
		http.Error(rw, "Error al recuperar el ID del cliente ", http.StatusInternalServerError)
		return
	}

	responseAllOrdersMenu, err := dbOrder.GetOrderViewer(client.ID)

	if err != nil {
		http.Error(rw, "Error a recuperar las ordenes de la BD", http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "aplication/json")
	rw.WriteHeader(http.StatusAccepted)
	json.NewEncoder(rw).Encode(responseAllOrdersMenu)

}
