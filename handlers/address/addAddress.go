package handlers

import (
	"encoding/json"
	"net/http"
	"viandasApp/dtos"

	dbaddress "viandasApp/db/address"
)

func AddAddress(rw http.ResponseWriter, r *http.Request) {

	var addressDto dtos.Address

	err := json.NewDecoder(r.Body).Decode(&addressDto)

	if err != nil {
		http.Error(rw, "Error en los datos recibidos "+err.Error(), 400)
		return
	}

	address := addressDto.ToModelAddress()

	address.Active = true
	address.CityID = 1

	status, err := dbaddress.AddAddress(*address, addressDto.IDClient)
	if err != nil {
		http.Error(rw, "Ocurrio un error al cargar la dirección "+err.Error(), http.StatusInternalServerError)
		return
	}

	if !status {
		http.Error(rw, "no se ha logrado cargar la dirección en la BD", http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-type", "application/json")
	rw.WriteHeader(http.StatusCreated)
}
