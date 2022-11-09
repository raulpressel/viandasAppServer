package handlers

import (
	"encoding/json"
	"net/http"
	"viandasApp/dtos"

	dbpathology "viandasApp/db/pathology"
)

func UploadPathology(rw http.ResponseWriter, r *http.Request) {

	var pathologyDto dtos.Pathology

	err := json.NewDecoder(r.Body).Decode(&pathologyDto)

	if err != nil {
		http.Error(rw, "Error en los datos recibidos "+err.Error(), 400)
		return
	}

	pathology := pathologyDto.ToModelPathology()

	status, err := dbpathology.UploadPathology(*pathology)
	if err != nil {
		http.Error(rw, "Ocurrio un error al cargar la patologia "+err.Error(), http.StatusInternalServerError)
		return
	}

	if !status {
		http.Error(rw, "no se ha logrado cargar la patologia en la BD", http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-type", "application/json")
	rw.WriteHeader(http.StatusCreated)

}
