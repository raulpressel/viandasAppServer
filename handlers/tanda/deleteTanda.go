package handlers

import (
	"net/http"
	"strconv"
	db "viandasApp/db/tanda"
)

func DeleteTanda(rw http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("idTanda")
	if len(ID) < 1 {
		http.Error(rw, "El parametro ID es obligatorio", http.StatusBadRequest)
		return
	}

	_ID, err := strconv.Atoi(ID)
	if err != nil {
		http.Error(rw, "Error al convertir el ID", http.StatusInternalServerError)
		return
	}

	tandaModel, err := db.GetTandaById(_ID)
	if err != nil {
		http.Error(rw, "No se pudo recuperar el plato de la base de datos "+err.Error(), http.StatusInternalServerError)
		return
	}

	tandaModel.Active = false

	status, err := db.DeleteTanda(tandaModel)
	if err != nil {
		http.Error(rw, "No se pudo borrar la tanda de la base de datos "+err.Error(), http.StatusInternalServerError)
		return
	}

	if !status {
		http.Error(rw, "No se pudo borrar la tanda de la base de datos ", http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-type", "application/json")
	rw.WriteHeader(http.StatusCreated)

}
