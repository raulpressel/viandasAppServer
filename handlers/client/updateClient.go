package handlers

import (
	"encoding/json"
	"net/http"
	"time"
	dbClient "viandasApp/db/client"
	"viandasApp/dtos"
	"viandasApp/models"
)

func UpdateClient(rw http.ResponseWriter, r *http.Request) {

	var clientDto dtos.RegisterRequest

	var clientModel models.Client

	var clientPathologyModel []models.ClientPathology

	var cPModel models.ClientPathology

	err := json.NewDecoder(r.Body).Decode(&clientDto)

	if err != nil {
		http.Error(rw, "Error en los datos recibidos "+err.Error(), http.StatusBadRequest)
		return
	}

	clientModel, err = dbClient.GetClientById(clientDto.Client.ID)

	if err != nil {
		http.Error(rw, "no fue posible recuperar el cliente en la BD", http.StatusInternalServerError)
		return
	}

	if clientModel.ID == 0 {
		http.Error(rw, "no fue posible recuperar el cliente en la BD", http.StatusBadRequest)
		return
	}

	clientModel.BornDate, err = time.Parse(time.RFC3339, clientDto.Client.BornDate)
	if err != nil {
		http.Error(rw, "Error en el formato de fecha recibido "+err.Error(), http.StatusBadRequest)
		return
	}

	if clientDto.Client.PhonePrimary == "" {
		http.Error(rw, "Debe cargar al menos un número teléfono "+err.Error(), http.StatusBadRequest)
		return
	}
	clientModel.PhonePrimary = clientDto.Client.PhonePrimary

	clientModel.Observation = clientDto.Client.ObsClient

	clientModel.PhoneSecondary = clientDto.Client.PhoneSecondary

	if len(clientDto.Client.Pathologies) > 0 {
		for _, path := range clientDto.Client.Pathologies {

			cPModel.PathologyID = path.ID

			clientPathologyModel = append(clientPathologyModel, cPModel)

		}

	}

	status, err := dbClient.UpdateClient(clientModel, clientPathologyModel)
	if err != nil {
		http.Error(rw, "Ocurrio un error al modificar el cliente "+err.Error(), http.StatusInternalServerError)
		return
	}

	if !status {
		http.Error(rw, "no fue posible modificar el cliente en la BD", http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-type", "application/json")
	rw.WriteHeader(http.StatusCreated)

}
