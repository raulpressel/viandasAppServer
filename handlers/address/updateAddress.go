package handlers

import (
	"encoding/json"
	"net/http"
	"viandasApp/dtos"
	"viandasApp/models"

	dbAddress "viandasApp/db/address"
)

func UpdateAddress(rw http.ResponseWriter, r *http.Request) {

	var addressDto dtos.Address

	var addressModel models.Address

	err := json.NewDecoder(r.Body).Decode(&addressDto)

	if err != nil {
		http.Error(rw, "Error en los datos recibidos "+err.Error(), 400)
		return
	}

	addressModel, err = dbAddress.GetAddressById(addressDto.Address.ID)

	addressModel.Street = addressDto.Address.Street
	addressModel.Number = addressDto.Address.Number
	addressModel.Floor = addressDto.Address.Floor
	addressModel.Departament = addressDto.Address.Departament
	addressModel.Observation = addressDto.Address.Observation

	status, err := dbAddress.UpdateAddress(addressModel)
	if err != nil {
		http.Error(rw, "Ocurrio un error al modificar la dirección "+err.Error(), http.StatusInternalServerError)
		return
	}

	if !status {
		http.Error(rw, "no se ha logrado modificar la dirección en la BD", http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-type", "application/json")
	rw.WriteHeader(http.StatusCreated)
}