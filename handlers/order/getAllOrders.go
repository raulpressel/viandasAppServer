package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	dbOrder "viandasApp/db/order"
	"viandasApp/dtos"
)

func GetAllOrders(rw http.ResponseWriter, r *http.Request) {

	dbOrder.FinishedOrder()

	var allOrderDto dtos.AllOrderRequest

	err := json.NewDecoder(r.Body).Decode(&allOrderDto)
	if err != nil {
		http.Error(rw, "Error en los datos recibidos "+err.Error(), http.StatusBadRequest)
		return
	}

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

	responseAllOrdersMenu, err := dbOrder.GetAllOrders(date, dateStart, dateEnd, allOrderDto.Active, allOrderDto.Cancel, allOrderDto.Finished, allOrderDto.Paid, allOrderDto.NotPaid)

	if err != nil {
		http.Error(rw, "Error a recuperar las ordenes de la BD", http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "aplication/json")
	rw.WriteHeader(http.StatusAccepted)
	json.NewEncoder(rw).Encode(responseAllOrdersMenu)

}
