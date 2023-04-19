package handlers

import (
	"encoding/json"
	"net/http"
	settingDB "viandasApp/db/setting"
	"viandasApp/dtos"
)

func UpdateZone(rw http.ResponseWriter, r *http.Request) {

	var zoneDto dtos.ZoneRequest

	err := json.NewDecoder(r.Body).Decode(&zoneDto)

	if err != nil {
		http.Error(rw, "Error en los datos recibidos "+err.Error(), http.StatusBadRequest)
		return
	}

	zoneModel, err := settingDB.GetZoneById(zoneDto.Zone.ID)

	if err != nil {
		http.Error(rw, "no fue posible recuperar la tanda de la BD", http.StatusInternalServerError)
		return
	}

	if zoneModel.ID == 0 {
		http.Error(rw, "no fue posible recuperar la tanda de la BD", http.StatusBadRequest)
		return
	}

	zoneModel.Description = zoneDto.Zone.Description

	if zoneDto.Zone.Price < 0 {
		http.Error(rw, "Error en los datos recibidos, no puede ser menor a 0 "+err.Error(), http.StatusBadRequest)
		return
	}

	zoneModel.Price = zoneDto.Zone.Price

	status, err := settingDB.UploadZone(zoneModel)

	if err != nil {
		http.Error(rw, "Ocurrio un error al intentar actualizar la zona "+err.Error(), http.StatusInternalServerError)
		return
	}

	if !status {
		http.Error(rw, "no se ha logrado actualizar la zona en la BD", http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-type", "application/json")
	rw.WriteHeader(http.StatusCreated)

}
