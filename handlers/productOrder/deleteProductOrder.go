package handlers

import (
	"net/http"
	"strconv"
	db "viandasApp/db/productOrder"
)

func DeleteProductOrder(rw http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("idProductOrder")
	if len(ID) < 1 {
		http.Error(rw, "debe enviar el parametro idProductOrder", http.StatusBadRequest)
		return
	}

	id, _ := strconv.Atoi(ID)
	ok, err := db.DeleteProductOrder(id)
	if err != nil || !ok {
		http.Error(rw, "No se pudo eliminar la orden de productos", http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusCreated)
}
