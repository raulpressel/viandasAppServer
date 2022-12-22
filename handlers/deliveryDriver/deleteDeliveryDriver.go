package handlers

import (
	"net/http"
	"strconv"
	deliDriver "viandasApp/db/deliveryDriver"
)

func DeleteDeliveryDriver(rw http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("idDeliveryDriver")

	if len(ID) < 1 {
		http.Error(rw, "debe enviar el parametro id", http.StatusBadRequest)
		return
	}

	idDeliveryDriver, err := strconv.Atoi(ID)

	if err != nil {
		http.Error(rw, "Error al convertir el ID", http.StatusInternalServerError)
		return
	}

	deliDriverModel, err := deliDriver.GetDeliveryDriverByID(idDeliveryDriver)

	if err != nil {
		http.Error(rw, "no fue posible recuperar el cadete en la BD", http.StatusInternalServerError)
		return
	}

	if deliDriverModel.ID == 0 {
		http.Error(rw, "no fue posible recuperar el cadete en la BD", http.StatusBadRequest)
		return
	}

	deliDriverModel.Active = false

	status, err := deliDriver.DeleteDeliveryDriver(deliDriverModel)

	if err != nil {
		http.Error(rw, "Ocurrio un error al eliminar el cadete "+err.Error(), http.StatusInternalServerError)
		return
	}

	if !status {
		http.Error(rw, "no se ha logrado eliminar el cadete en la BD", http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-type", "application/json")
	rw.WriteHeader(http.StatusCreated)

}
