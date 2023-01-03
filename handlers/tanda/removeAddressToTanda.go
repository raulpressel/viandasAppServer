package handlers

import (
	"encoding/json"
	"net/http"
	addrDb "viandasApp/db/address"
	tandaDb "viandasApp/db/tanda"
	"viandasApp/dtos"
	"viandasApp/models"
)

func RemoveAddressToTanda(rw http.ResponseWriter, r *http.Request) {

	var tandaAddDto dtos.TandaAddressRequest

	var tandaAddressModel models.TandaAddress

	var tandaAddressesModel []models.TandaAddress

	err := json.NewDecoder(r.Body).Decode(&tandaAddDto)

	if err != nil {
		http.Error(rw, "Error en los datos recibidos "+err.Error(), http.StatusBadRequest)
		return
	}

	tandaModel, err := tandaDb.GetTandaById(tandaAddDto.IDTanda)

	if tandaModel.ID == 0 {
		http.Error(rw, "No existe la tanda solicitada "+err.Error(), http.StatusBadRequest)
		return
	}

	for _, val := range tandaAddDto.IDAddress {
		addrModel, err := addrDb.GetAddressById(val)
		if addrModel.ID == 0 {
			http.Error(rw, "No existe la direccion solicitada "+err.Error(), http.StatusBadRequest)
			return
		}
		if err != nil {
			http.Error(rw, "Error al recuperar el modelo de direccion de la BD "+err.Error(), http.StatusInternalServerError)
			return
		}

		tandaAddressModel.TandaID = tandaAddDto.IDTanda
		tandaAddressModel.AddressID = val
		tandaAddressesModel = append(tandaAddressesModel, tandaAddressModel)

	}

	status, err := tandaDb.RemoveAddressToTanda(tandaAddressesModel)

	if err != nil {
		http.Error(rw, "Ocurrio un error al intentar quitar la o las direcciones de la tanda "+err.Error(), http.StatusInternalServerError)
		return
	}

	if !status {
		http.Error(rw, "no se ha logrado quitar la o las direcciones de la tanda en la BD", http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-type", "application/json")
	rw.WriteHeader(http.StatusCreated)

}
