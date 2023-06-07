package handlers

import (
	"encoding/json"
	"net/http"
	"regexp"
	"time"
	dbClient "viandasApp/db/client"
	"viandasApp/dtos"
	"viandasApp/models"
)

func AddClient(rw http.ResponseWriter, r *http.Request) {

	var registerDto dtos.RegisterRequest

	var clientModel models.Client

	var addressModel []models.Address

	var addModel models.Address

	var clientPathologyModel []models.ClientPathology

	var cPModel models.ClientPathology

	err := json.NewDecoder(r.Body).Decode(&registerDto)

	if err != nil {
		http.Error(rw, "Error en los datos recibidos "+err.Error(), http.StatusBadRequest)
		return
	}

	clientModel.Name = registerDto.Client.Name
	clientModel.LastName = registerDto.Client.LastName
	clientModel.Email = registerDto.Client.Email //ver si validar que sea mail

	valid := isValidEmail(clientModel.Email)
	if !valid {
		http.Error(rw, "El email ingresado no es válido ", http.StatusBadRequest)
	}

	cli, err := dbClient.GetClientByEmail(clientModel.Email)
	if err != nil {
		http.Error(rw, "No fue posible recuperar el cliente con el correo electronico "+err.Error(), http.StatusBadRequest)
		return
	}
	if cli.ID > 0 {
		http.Error(rw, "El email ingresado ya se encuentra asociado a un cliente ", http.StatusBadRequest)
	}

	clientModel.Active = true

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
			addModel.IDZone = addr.IDZone
			addModel.Lat = addr.Lat
			addModel.Lng = addr.Lng
			addModel.CityID = 1
			addModel.Favourite = true
			addModel.Active = true

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

func isValidEmail(email string) bool {
	// Expresión regular para verificar el formato del correo electrónico
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, err := regexp.MatchString(regex, email)
	if err != nil {
		return false
	}
	return match
}
