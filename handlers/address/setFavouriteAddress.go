package handlers

import (
	"encoding/json"
	"net/http"
	dbAddress "viandasApp/db/address"
	"viandasApp/dtos"
	"viandasApp/models"
)

func SetFavouriteAddress(rw http.ResponseWriter, r *http.Request) {
	var addressDto dtos.Address
	var oldFavAddressModel models.Address
	var newFavAddressModel models.Address

	err := json.NewDecoder(r.Body).Decode(&addressDto)

	if err != nil {
		http.Error(rw, "Error en los datos recibidos "+err.Error(), 400)
		return
	}

	oldFavAddressModel, err = dbAddress.GetAddressById(addressDto.IDOldFavouriteAddress)
	oldFavAddressModel.Favourite = false

	newFavAddressModel, err = dbAddress.GetAddressById(addressDto.IDNewFavouriteAddress)
	newFavAddressModel.Favourite = true

	statusOld, err := dbAddress.UpdateAddress(oldFavAddressModel)
	if err != nil {
		http.Error(rw, "Ocurrio un error al modificar la dirección "+err.Error(), http.StatusInternalServerError)
		return
	}

	if !statusOld {
		http.Error(rw, "no se ha logrado modificar la dirección en la BD", http.StatusInternalServerError)
		return
	}

	statusNew, err := dbAddress.UpdateAddress(newFavAddressModel)
	if err != nil {
		http.Error(rw, "Ocurrio un error al modificar la dirección "+err.Error(), http.StatusInternalServerError)
		return
	}

	if !statusNew {
		http.Error(rw, "no se ha logrado modificar la dirección en la BD", http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-type", "application/json")
	rw.WriteHeader(http.StatusCreated)
}
