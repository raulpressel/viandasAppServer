package handlers

import (
	"net/http"
	"strconv"
	db "viandasApp/db/productOrder"
)

func UpdateProductOrderStatus(rw http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("idProductOrder")
	status := r.URL.Query().Get("status")

	if len(ID) < 1 {
		http.Error(rw, "debe enviar el parametro idProductOrder", http.StatusBadRequest)
		return
	}
	if status != "pendiente" && status != "entregado" {
		http.Error(rw, "status debe ser 'pendiente' o 'entregado'", http.StatusBadRequest)
		return
	}

	id, _ := strconv.Atoi(ID)
	model, err := db.GetProductOrderById(id)
	if err != nil {
		http.Error(rw, "no fue posible recuperar la orden por ID", http.StatusInternalServerError)
		return
	}

	model.Status = status

	ok, err := db.UpdateProductOrderStatus(model)
	if err != nil || !ok {
		http.Error(rw, "No se pudo actualizar el estado de la orden", http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusCreated)
}
