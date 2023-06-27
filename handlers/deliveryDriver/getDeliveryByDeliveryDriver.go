package handlers

import (
	"encoding/json"
	"net/http"
	"time"
	deliDriver "viandasApp/db/deliveryDriver"
	"viandasApp/dtos"
	"viandasApp/models"
)

func GetDeliveryByDeliveryDriver(rw http.ResponseWriter, r *http.Request) {

	var deliveryDto dtos.DeliveryRequest

	var deliveryDriverModel models.DeliveryDriver

	err := json.NewDecoder(r.Body).Decode(&deliveryDto)

	if err != nil {
		http.Error(rw, "Error en los datos recibidos "+err.Error(), http.StatusBadRequest)
		return
	}

	if deliveryDto.DeliveryDriverID != nil {
		deliveryDriverModel, err = deliDriver.GetDeliveryDriverByID(*deliveryDto.DeliveryDriverID)
		if err != nil {
			http.Error(rw, "Error al obtener el delivery driver", http.StatusInternalServerError)
			return
		}
		if deliveryDriverModel.ID == 0 {
			http.Error(rw, "No existe un Delivery Driver con ese ID ", http.StatusBadRequest)
			return
		}
	}

	dateStart, err := time.Parse(time.RFC3339, deliveryDto.DateStart)
	if err != nil {
		http.Error(rw, "Error en el formato de fecha recibido "+err.Error(), http.StatusBadRequest)
		return
	}

	dateEnd, err := time.Parse(time.RFC3339, deliveryDto.DateEnd)
	if err != nil {
		http.Error(rw, "Error en el formato de fecha recibido "+err.Error(), http.StatusBadRequest)
		return
	}

	response, err := deliDriver.GetDeliveryByDeliveryDriver(deliveryDriverModel.ID, dateStart, dateEnd)

	if err != nil {
		http.Error(rw, "Ocurrio un error al intentar registrar el cadete "+err.Error(), http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "aplication/json")
	rw.WriteHeader(http.StatusAccepted)
	json.NewEncoder(rw).Encode(response)
}
