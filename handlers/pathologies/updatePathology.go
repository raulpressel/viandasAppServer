package handlers

import (
	"encoding/json"
	"net/http"

	dbpathology "viandasApp/db/pathology"

	"viandasApp/dtos"
	"viandasApp/models"
)

func UpdatePathology(rw http.ResponseWriter, r *http.Request) {

	rw.Header().Add("content-type", "application/json")

	var pathologyDto dtos.Pathology

	err := json.NewDecoder(r.Body).Decode(&pathologyDto)

	if err != nil {
		http.Error(rw, "Error en los datos recibidos "+err.Error(), http.StatusBadRequest)
		return
	}

	var pathologyModel models.Pathology

	pathologyModel, err = dbpathology.GetPathologyById(pathologyDto.Pathology.ID)

	if err != nil {
		http.Error(rw, "Error al recuperar la patologia de la BD "+err.Error(), http.StatusInternalServerError)
		return
	}

	if pathologyDto.Pathology.Description != pathologyModel.Description {
		pathologyModel.Description = pathologyDto.Pathology.Description
	}

	if pathologyDto.Pathology.Color != pathologyModel.Color {
		pathologyModel.Color = pathologyDto.Pathology.Color
	}

	pathologyModel.Active = true

	status, err := dbpathology.UpdatePathology(pathologyModel)

	if err != nil {
		http.Error(rw, "Ocurrio un error al cargar la patologia "+err.Error(), http.StatusInternalServerError)
		return
	}

	if !status {
		http.Error(rw, "no se ha logrado modificar la patologia en la BD", http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-type", "application/json")
	rw.WriteHeader(http.StatusCreated)

}
