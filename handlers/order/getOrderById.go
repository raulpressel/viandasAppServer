package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	db "viandasApp/db/order"
)

func GetOrderById(rw http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("idOrder")

	if len(ID) < 1 {
		http.Error(rw, "debe enviar el parametro id", http.StatusBadRequest)
		return
	}

	idOrder, err := strconv.Atoi(ID)

	if err != nil {
		http.Error(rw, "Error al convertir el ID", http.StatusInternalServerError)
		return
	}

	responseOrder, err := db.GetOrderById(idOrder)

	if err != nil {
		http.Error(rw, "Orden no encontrada", http.StatusBadRequest)
		return
	}

	if responseOrder.ID == 0 {
		http.Error(rw, "No hay orden en la BD", http.StatusNotFound)
		return
	}

	rw.Header().Set("Content-Type", "aplication/json")
	rw.WriteHeader(http.StatusAccepted)
	json.NewEncoder(rw).Encode(responseOrder)

}
