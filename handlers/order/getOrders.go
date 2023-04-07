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

	var dat daterr
	err := json.NewDecoder(r.Body).Decode(&dat)
	if err != nil {
		http.Error(rw, "Error en los datos recibidos "+err.Error(), http.StatusBadRequest)
		return
	}

	date, err := time.Parse(time.RFC3339, dat.Date)
	if err != nil {
		http.Error(rw, "Error en el formato de fecha recibido "+err.Error(), http.StatusBadRequest)
		return
	}

	responseAllOrdersMenu, err := dbOrder.GetOrders(date)

	if err != nil {
		http.Error(rw, "Error a recuperar las ordenes de la BD", http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "aplication/json")
	rw.WriteHeader(http.StatusAccepted)
	json.NewEncoder(rw).Encode(responseAllOrdersMenu)

}
