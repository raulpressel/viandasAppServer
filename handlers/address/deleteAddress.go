package handlers

import (
	"net/http"
	"strconv"
	dbAddress "viandasApp/db/address"
)

func DeleteAddress(rw http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("idAddress")

	if len(ID) < 1 {
		http.Error(rw, "debe enviar el parametro id", http.StatusBadRequest)
		return
	}

	idAddress, err := strconv.Atoi(ID)

	if err != nil {
		http.Error(rw, "Error al convertir el ID", http.StatusInternalServerError)
		return
	}

	addressModel, err := dbAddress.GetAddressById(idAddress)

	if err != nil {
		http.Error(rw, "Error al recuperar el modelo de direccion de la BD "+err.Error(), http.StatusInternalServerError)
		return
	}

	addressModel.Active = false

	status, err := dbAddress.DeleteAddress(addressModel)

	if err != nil {
		http.Error(rw, "Ocurrio un error al eliminar la direccion "+err.Error(), http.StatusInternalServerError)
		return
	}

	if !status {
		http.Error(rw, "no se ha logrado eliminar la direccion en la BD", http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-type", "application/json")
	rw.WriteHeader(http.StatusCreated)

}
