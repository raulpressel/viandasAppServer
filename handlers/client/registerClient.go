package handlers

import (
	"encoding/json"
	"net/http"
	"time"
	"viandasApp/db"
	dbClient "viandasApp/db/client"
	"viandasApp/dtos"
	"viandasApp/handlers"
	"viandasApp/models"
)

func RegisterClient(rw http.ResponseWriter, r *http.Request) {

	var clientDto dtos.Client

	var clientModel models.Client

	var addressModel []models.Address

	var addModel models.Address

	var clientPathologyModel []models.ClientPathology

	var cPModel models.ClientPathology

	var addcliModel models.ClientAddress

	db.ExistTable(clientModel)
	db.ExistTable(addModel)
	db.ExistTable(cPModel)
	db.ExistTable(addcliModel)

	cli := handlers.GetClient()

	clientModel.IDUserKL = cli.ID
	clientModel.Name = cli.Name
	clientModel.LastName = cli.LastName
	clientModel.Email = cli.Email

	cm, res := dbClient.CheckExistClient(clientModel.IDUserKL)

	if res {
		http.Error(rw, "Ya existe un cliente con los datos solicitados ", http.StatusBadRequest)
		return
	}

	if cm.ID > 0 {
		http.Error(rw, "Ya existe un cliente con los datos solicitados ", http.StatusBadRequest)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&clientDto)

	if err != nil {
		http.Error(rw, "Error en los datos recibidos "+err.Error(), http.StatusBadRequest)
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

	if len(clientDto.Client.Address) > 0 {

		for _, addr := range clientDto.Client.Address {

			addModel.Street = addr.Street
			addModel.Number = addr.Number
			addModel.Floor = addr.Floor
			addModel.Departament = addr.Departament
			addModel.Observation = addr.Observation
			addModel.CityID = 1

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
