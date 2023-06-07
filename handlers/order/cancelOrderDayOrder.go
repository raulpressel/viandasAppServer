package handlers

import (
	"net/http"
	"strconv"
	db "viandasApp/db/order"
)

func CancelOrderDayOrder(rw http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("idDayOrder")

	if len(ID) < 1 {
		http.Error(rw, "debe enviar el parametro id", http.StatusBadRequest)
		return
	}

	idOrder, err := strconv.Atoi(ID)

	if err != nil {
		http.Error(rw, "Error al convertir el ID", http.StatusInternalServerError)
		return
	}

	modelDayOrder, err := db.GetDayOrderById(idOrder)

	if err != nil {
		http.Error(rw, "Day Orden no encontrado con el ID solicitado", http.StatusBadRequest)
		return
	}

	modelDayOrder.Status = false

	status, err := db.CancelDayOrder(modelDayOrder)

	if err != nil {
		http.Error(rw, "error al actualizar el estado del DayOrder", http.StatusInternalServerError)
		return
	}

	if !status {
		http.Error(rw, "no se pudo cancelar la orden", http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-type", "application/json")
	rw.WriteHeader(http.StatusCreated)

}
