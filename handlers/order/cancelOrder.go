package handlers

import (
	"net/http"
	"strconv"
	db "viandasApp/db/order"
)

func CancelOrder(rw http.ResponseWriter, r *http.Request) {

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

	modelOrder, err := db.GetModelOrderById(idOrder)

	if err != nil {
		http.Error(rw, "Orden no encontrada", http.StatusBadRequest)
		return
	}

	if modelOrder.ID == 0 {
		http.Error(rw, "No hay orden en la BD", http.StatusNotFound)
		return
	}

	modelOrder.Status = false

	status, err := db.CancelOrder(modelOrder)

	if err != nil {
		http.Error(rw, "Orden no encontrada", http.StatusInternalServerError)
		return
	}

	if !status {
		http.Error(rw, "no se pudo cancelar la orden", http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-type", "application/json")
	rw.WriteHeader(http.StatusCreated)

}
