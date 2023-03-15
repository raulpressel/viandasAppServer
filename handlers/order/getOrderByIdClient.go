package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	dbOrder "viandasApp/db/order"
)

func getOrderByIdClient(rw http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("idClient")

	if len(id) < 1 {
		http.Error(rw, "debe enviar el parametro ID", http.StatusBadRequest)
		return
	}

	idClient, err := strconv.Atoi(id)

	if err != nil {
		http.Error(rw, "Error al convertir el ID", http.StatusInternalServerError)
		return
	}

	responseAllOrdersMenu, err := dbOrder.GetAllOrder(idClient)

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
