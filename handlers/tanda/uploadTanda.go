package handlers

import (
	"encoding/json"
	"net/http"
	"viandasApp/db"
	deliDriver "viandasApp/db/deliveryDriver"
	tandaDb "viandasApp/db/tanda"
	"viandasApp/dtos"
	"viandasApp/models"
)

func UploadTanda(rw http.ResponseWriter, r *http.Request) {

	var tandaDto dtos.TandaRequest

	var tandaModel models.Tanda

	err := json.NewDecoder(r.Body).Decode(&tandaDto)

	if err != nil {
		http.Error(rw, "Error en los datos recibidos "+err.Error(), http.StatusBadRequest)
		return
	}

	db.ExistTable(tandaModel)

	deliDriverModel, err := deliDriver.GetDeliveryDriverByID(tandaDto.IdDeliveryDriver)

	if err != nil {
		http.Error(rw, "no fue posible recuperar el cadete en la BD", http.StatusInternalServerError)
		return
	}

	if deliDriverModel.ID == 0 {
		http.Error(rw, "no fue posible recuperar el cadete en la BD", http.StatusBadRequest)
		return
	}

	tandaModel.Description = tandaDto.Tanda.Description

	tandaModel.DeliveryDriverID = deliDriverModel.ID

	tandaModel.HourStart = tandaDto.Tanda.HourStart
	tandaModel.HourEnd = tandaDto.Tanda.HourEnd

	tandaModel.DeliveryDriverID = deliDriverModel.ID

	tandaModel.Active = true

	status, err := tandaDb.UploadTanda(tandaModel)

	if err != nil {
		http.Error(rw, "Ocurrio un error al intentar registrar la tanda "+err.Error(), http.StatusInternalServerError)
		return
	}

	if !status {
		http.Error(rw, "no se ha logrado registrar la tanda en la BD", http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-type", "application/json")
	rw.WriteHeader(http.StatusCreated)

}
