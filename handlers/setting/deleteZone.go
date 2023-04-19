package handlers

import (
	"net/http"
	"strconv"
	db "viandasApp/db/setting"
)

func DeleteZone(rw http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("idZone")
	if len(ID) < 1 {
		http.Error(rw, "El parametro ID es obligatorio", http.StatusBadRequest)
		return
	}

	_ID, err := strconv.Atoi(ID)
	if err != nil {
		http.Error(rw, "Error al convertir el ID", http.StatusInternalServerError)
		return
	}

	zoneModel, err := db.GetZoneById(_ID)
	if err != nil {
		http.Error(rw, "No se pudo recuperar la zona de la base de datos "+err.Error(), http.StatusInternalServerError)
		return
	}

	zoneModel.Active = false

	status, err := db.DeleteZone(zoneModel)

	if err != nil {
		http.Error(rw, "No se pudo borrar la zona de la base de datos "+err.Error(), http.StatusInternalServerError)
		return
	}

	if !status {
		http.Error(rw, "No se pudo borrar la zona de la base de datos ", http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-type", "application/json")
	rw.WriteHeader(http.StatusCreated)

}
