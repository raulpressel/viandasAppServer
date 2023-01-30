package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	dbOrder "viandasApp/db/order"
)

type daterr struct {
	Date string `json:"date"`
}

func GetOrders(rw http.ResponseWriter, r *http.Request) {

	/* usr := handlers.GetUser()

	client, valid := dbClient.CheckExistClient(usr.ID)

	if !valid {
		http.Error(rw, "Error al recuperar el ID del cliente ", http.StatusInternalServerError)
		return
	} */

	var dat daterr
	err := json.NewDecoder(r.Body).Decode(&dat)

	//	var responseAllOrdersMenu dtos.OrderViewerResponse

	test, err := time.Parse(time.RFC3339, dat.Date)
	if err != nil {
		http.Error(rw, "Error en el formato de fecha recibido "+err.Error(), http.StatusBadRequest)
		return
	}

	responseAllOrdersMenu, err := dbOrder.GetOrders(test)

	if err != nil {
		http.Error(rw, "Error a recuperar las ordenes de la BD", http.StatusInternalServerError)
		return
	}

	/* if responseAllOrdersMenu.Order == nil {
		http.Error(rw, "No hay ordenes en la BD", http.StatusNotFound)
	} */

	rw.Header().Set("Content-Type", "aplication/json")
	rw.WriteHeader(http.StatusAccepted)
	json.NewEncoder(rw).Encode(responseAllOrdersMenu)

}
