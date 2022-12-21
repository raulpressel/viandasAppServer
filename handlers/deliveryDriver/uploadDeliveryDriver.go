package handlers

import (
	"encoding/json"
	"net/http"
	"time"
	"viandasApp/db"
	deliDriver "viandasApp/db/deliveryDriver"
	"viandasApp/dtos"
	"viandasApp/models"
)

func UploadDeliveryDriver(rw http.ResponseWriter, r *http.Request) {

	var deliDriverDto dtos.DeliveryDriverRequest

	var deliDriverModel models.DeliveryDriver

	var vehicleModel models.Vehicle

	var addressModel models.Address

	err := json.NewDecoder(r.Body).Decode(&deliDriverDto)

	if err != nil {
		http.Error(rw, "Error en los datos recibidos "+err.Error(), http.StatusBadRequest)
		return
	}

	db.ExistTable(deliDriverModel)

	db.ExistTable(vehicleModel)

	deliDriverModel.DNI = deliDriverDto.DeliveryDriver.DNI

	deliDriverModel.Name = deliDriverDto.DeliveryDriver.Name

	deliDriverModel.LastName = deliDriverDto.DeliveryDriver.LastName

	deliDriverModel.Phone = deliDriverDto.DeliveryDriver.Phone

	deliDriverModel.BornDate, err = time.Parse(time.RFC3339, deliDriverDto.DeliveryDriver.BornDate)
	if err != nil {
		http.Error(rw, "Error en el formato de fecha recibido "+err.Error(), http.StatusBadRequest)
		return
	}

	deliDriverModel.Active = true

	vehicleModel.Brand = deliDriverDto.DeliveryDriver.Vehicle.Brand
	vehicleModel.Models = deliDriverDto.DeliveryDriver.Vehicle.Models
	vehicleModel.Patent = deliDriverDto.DeliveryDriver.Vehicle.Patent
	vehicleModel.Year = deliDriverDto.DeliveryDriver.Vehicle.Year

	addressModel.Street = deliDriverDto.DeliveryDriver.Address.Street
	addressModel.Number = deliDriverDto.DeliveryDriver.Address.Number
	addressModel.Floor = deliDriverDto.DeliveryDriver.Address.Floor
	addressModel.Departament = deliDriverDto.DeliveryDriver.Address.Departament
	addressModel.Observation = deliDriverDto.DeliveryDriver.Address.Observation

	status, err := deliDriver.UploadDeliveryDriver(deliDriverModel, vehicleModel, addressModel)

	if err != nil {
		http.Error(rw, "Ocurrio un error al intentar registrar el cadete "+err.Error(), http.StatusInternalServerError)
		return
	}

	if !status {
		http.Error(rw, "no se ha logrado registrar el cadete en la BD", http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-type", "application/json")
	rw.WriteHeader(http.StatusCreated)

}
