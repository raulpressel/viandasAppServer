package handlers

import (
	"encoding/json"
	"net/http"
	dbAddress "viandasApp/db/address"
	"viandasApp/dtos"
	"viandasApp/models"
)

func AddressTakeAway(rw http.ResponseWriter, r *http.Request) {

	var addressDto dtos.AddressRequest
	var addressModel models.Address

	err := json.NewDecoder(r.Body).Decode(&addressDto)

	if err != nil {
		http.Error(rw, "Error en los datos recibidos "+err.Error(), http.StatusBadRequest)
		return
	}

	if addressDto.ID > 0 {
		addressModel, err = dbAddress.GetAddressById(addressDto.ID)
		if err != nil {
			http.Error(rw, "Error al recuperar el modelo de direccion de la BD "+err.Error(), http.StatusInternalServerError)
			return
		}
	}
	addressModel.ID = 100
	addressModel.Street = addressDto.Street
	addressModel.Floor = addressDto.Floor
	addressModel.Number = addressDto.Number
	addressModel.Departament = addressDto.Departament
	addressModel.Observation = addressDto.Observation
	addressModel.IDZone = addressDto.IDZone
	addressModel.Lat = addressDto.Lat
	addressModel.Lng = addressDto.Lng
	addressModel.Active = true
	addressModel.CityID = 1

	status, err := dbAddress.AddressTakeAway(addressModel)

	if err != nil {
		http.Error(rw, "Ocurrio un error al intentar agregar/actualizar la dirección "+err.Error(), http.StatusInternalServerError)
		return
	}

	if !status {
		http.Error(rw, "no se ha logrado agregar/actualizar la dirección en la BD", http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-type", "application/json")
	rw.WriteHeader(http.StatusCreated)

}
