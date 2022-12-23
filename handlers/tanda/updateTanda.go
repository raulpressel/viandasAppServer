package handlers

import (
	"encoding/json"
	"net/http"
	deliDriver "viandasApp/db/deliveryDriver"
	tandaDb "viandasApp/db/tanda"
	"viandasApp/dtos"
)

func UpdateTanda(rw http.ResponseWriter, r *http.Request) {

	var tandaDto dtos.TandaRequest

	err := json.NewDecoder(r.Body).Decode(&tandaDto)

	if err != nil {
		http.Error(rw, "Error en los datos recibidos "+err.Error(), http.StatusBadRequest)
		return
	}

	tandaModel, err := tandaDb.GetTandaById(tandaDto.Tanda.ID)

	if err != nil {
		http.Error(rw, "no fue posible recuperar la tanda de la BD", http.StatusInternalServerError)
		return
	}

	if tandaModel.ID == 0 {
		http.Error(rw, "no fue posible recuperar la tanda de la BD", http.StatusBadRequest)
		return
	}

	deliDriverModel, err := deliDriver.GetDeliveryDriverByID(tandaDto.IdDeliveryDriver)

	if err != nil {
		http.Error(rw, "no fue posible recuperar el cadete en la BD", http.StatusInternalServerError)
		return
	}

	if deliDriverModel.ID == 0 {
		http.Error(rw, "no fue posible recuperar el cadete en la BD", http.StatusBadRequest)
		return
	}

	tandaModel.DeliveryDriverID = deliDriverModel.ID

	tandaModel.Description = tandaDto.Tanda.Description

	tandaModel.HourStart = tandaDto.Tanda.HourStart
	tandaModel.HourEnd = tandaDto.Tanda.HourEnd

	tandaModel.Active = true

	status, err := tandaDb.UploadTanda(tandaModel)

	if err != nil {
		http.Error(rw, "Ocurrio un error al intentar actualizar la tanda "+err.Error(), http.StatusInternalServerError)
		return
	}

	if !status {
		http.Error(rw, "no se ha logrado actualizar la tanda en la BD", http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-type", "application/json")
	rw.WriteHeader(http.StatusCreated)

}
