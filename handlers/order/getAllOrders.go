package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	dbOrder "viandasApp/db/order"
	"viandasApp/dtos"
)

func GetAllOrders(rw http.ResponseWriter, r *http.Request) {

	var allOrderDto dtos.AllOrderRequest

	err := json.NewDecoder(r.Body).Decode(&allOrderDto)

	var dateStart time.Time

	var dateEnd time.Time

	date := false

	if allOrderDto.DateStart != nil && allOrderDto.DateEnd != nil {

		dateStart, err = time.Parse(time.RFC3339, *allOrderDto.DateStart)
		if err != nil {
			http.Error(rw, "Error en el formato de fecha recibido "+err.Error(), http.StatusBadRequest)
			return
		}

		dateEnd, err = time.Parse(time.RFC3339, *allOrderDto.DateEnd)
		if err != nil {
			http.Error(rw, "Error en el formato de fecha recibido "+err.Error(), http.StatusBadRequest)
			return
		}

		date = true

	}

	if err != nil {
		http.Error(rw, "Error a recuperar las ordenes de la BD", http.StatusInternalServerError)
		return
	}

	if err != nil {
		http.Error(rw, "Error a recuperar las ordenes de la BD", http.StatusInternalServerError)
		return
	}

	responseAllOrdersMenu, err := dbOrder.GetAllOrders(date, dateStart, dateEnd, allOrderDto.Active, allOrderDto.Inactive, allOrderDto.Paid, allOrderDto.NotPaid)

	rw.Header().Set("Content-Type", "aplication/json")
	rw.WriteHeader(http.StatusAccepted)
	json.NewEncoder(rw).Encode(responseAllOrdersMenu)

}
