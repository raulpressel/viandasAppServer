package handlers

import (
	"net/http"
	"strconv"
	db "viandasApp/db/setting"
)

func DeleteDiscount(rw http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("idDiscount")
	if len(ID) < 1 {
		http.Error(rw, "El parametro ID es obligatorio", http.StatusBadRequest)
		return
	}

	_ID, err := strconv.Atoi(ID)
	if err != nil {
		http.Error(rw, "Error al convertir el ID", http.StatusInternalServerError)
		return
	}

	discountModel, err := db.GetDiscountById(_ID)
	if err != nil {
		http.Error(rw, "No se pudo recuperar el descuento de la base de datos "+err.Error(), http.StatusInternalServerError)
		return
	}

	discountModel.Active = false

	status, err := db.DeleteDiscount(discountModel)

	if err != nil {
		http.Error(rw, "No se pudo borrar el descuento de la base de datos "+err.Error(), http.StatusInternalServerError)
		return
	}

	if !status {
		http.Error(rw, "No se pudo borrar el descuento de la base de datos ", http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-type", "application/json")
	rw.WriteHeader(http.StatusCreated)

}
