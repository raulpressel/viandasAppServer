package handlers

import (
	"net/http"
	"strconv"
	db "viandasApp/db/order"
)

func UpdateDayOrderAddress(rw http.ResponseWriter, r *http.Request) {

	idDO := r.URL.Query().Get("idDayOrder")

	if len(idDO) < 1 {
		http.Error(rw, "El parametro idDayOrder es obligatorio", http.StatusBadRequest)
		return
	}

	idA := r.URL.Query().Get("idAddress")

	if len(idA) < 1 {
		http.Error(rw, "El parametro idAddress es obligatorio", http.StatusBadRequest)
		return
	}

	idDayOrder, err := strconv.Atoi(idDO)

	if err != nil {
		http.Error(rw, "Error al convertir el ID", http.StatusInternalServerError)
		return
	}

	idAddress, err := strconv.Atoi(idA)

	if err != nil {
		http.Error(rw, "Error al convertir el ID", http.StatusInternalServerError)
		return
	}

	status, err := db.UpdateDayOrderAddress(idDayOrder, idAddress)

	if err != nil {
		http.Error(rw, "Orden no encontrada", http.StatusBadRequest)
		return
	}

	if !status {
		http.Error(rw, "no se ha logrado modificar la direccion de la orden", http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-type", "application/json")
	rw.WriteHeader(http.StatusCreated)

}
