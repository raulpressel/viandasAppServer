package handlers

import (
	"encoding/json"
	"net/http"
	"time"
	dbClient "viandasApp/db/client"
	"viandasApp/dtos"
	"viandasApp/handlers"
	"viandasApp/models"
)

func RegisterClient(rw http.ResponseWriter, r *http.Request) {

	var registerDto dtos.RegisterRequest

	var clientModel models.Client

	var addressModel []models.Address

	var addModel models.Address

	var clientPathologyModel []models.ClientPathology

	var cPModel models.ClientPathology

	usr := handlers.GetUser()

	clientModel.IDUserKL = usr.ID
	clientModel.Name = usr.Name
	clientModel.LastName = usr.LastName
	clientModel.Email = usr.Email

	cm, res := dbClient.CheckExistClient(clientModel.IDUserKL)

	if res {
		http.Error(rw, "Ya existe un cliente con los datos solicitados ", http.StatusBadRequest)
		return
	}

	if cm.ID > 0 {
		http.Error(rw, "Ya existe un cliente con los datos solicitados ", http.StatusBadRequest)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&registerDto)

	if err != nil {
		http.Error(rw, "Error en los datos recibidos "+err.Error(), http.StatusBadRequest)
		return
	}

	clientModel.BornDate, err = time.Parse(time.RFC3339, registerDto.Client.BornDate)
	if err != nil {
		http.Error(rw, "Error en el formato de fecha recibido "+err.Error(), http.StatusBadRequest)
		return
	}

	if registerDto.Client.PhonePrimary == "" {
		http.Error(rw, "Debe cargar al menos un número teléfono "+err.Error(), http.StatusBadRequest)
		return
	}
	clientModel.PhonePrimary = registerDto.Client.PhonePrimary

	clientModel.Observation = registerDto.Client.ObsClient

	clientModel.PhoneSecondary = registerDto.Client.PhoneSecondary

	if len(registerDto.Client.Pathologies) > 0 {
		for _, path := range registerDto.Client.Pathologies {

			cPModel.PathologyID = path.ID

			clientPathologyModel = append(clientPathologyModel, cPModel)

		}

	}

	if len(registerDto.Client.Address) > 0 {

		for _, addr := range registerDto.Client.Address {

			addModel.Street = addr.Street
			addModel.Number = addr.Number
			addModel.Floor = addr.Floor
			addModel.Departament = addr.Departament
			addModel.Observation = addr.Observation
			addModel.CityID = 1
			addModel.Favourite = true

			addressModel = append(addressModel, addModel)

		}
	} else {
		http.Error(rw, "Debe cargar al menos una dirección "+err.Error(), http.StatusBadRequest)
		return
	}

	status, err := dbClient.RegisterClient(clientModel, clientPathologyModel, addressModel)
	if err != nil {
		http.Error(rw, "Ocurrio un error al registrar el cliente "+err.Error(), http.StatusInternalServerError)
		return
	}

	if !status {
		http.Error(rw, "no fue posible registrar el cliente en la BD", http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-type", "application/json")
	rw.WriteHeader(http.StatusCreated)

}
