package handlers

import (
	"encoding/json"
	"net/http"
	"time"
	dbMenu "viandasApp/db/menu"
	"viandasApp/dtos"
)

func ValidateDateMenu(rw http.ResponseWriter, r *http.Request) {

	var validateDateMenuRequestDto dtos.ValidateDateMenuRequest

	var validateDateMenuResponseDto dtos.ValidateDateMenuRespone

	err := json.NewDecoder(r.Body).Decode(&validateDateMenuRequestDto)

	if err != nil {
		http.Error(rw, "Error en los datos recibidos "+err.Error(), http.StatusBadRequest)
		return
	}

	dateStart, err := time.Parse(time.RFC3339, validateDateMenuRequestDto.DateStart)
	if err != nil {
		http.Error(rw, "Error en el formato de fecha recibido "+err.Error(), http.StatusBadRequest)
		return
	}
	dateEnd, err := time.Parse(time.RFC3339, validateDateMenuRequestDto.DateEnd)
	if err != nil {
		http.Error(rw, "Error en el formato de fecha recibido "+err.Error(), http.StatusBadRequest)
		return
	}

	id, err := dbMenu.GetIdMenuActiveByDate(dateStart, dateEnd)
	if err != nil {
		http.Error(rw, "Ocurrio un error en la BD "+err.Error(), http.StatusInternalServerError)
		return
	}

	if id == 0 {
		validateDateMenuResponseDto.ValidDateMenu = true
	}

	rw.Header().Set("Content-Type", "aplication/json")
	rw.WriteHeader(http.StatusAccepted)
	json.NewEncoder(rw).Encode(validateDateMenuResponseDto)

}
