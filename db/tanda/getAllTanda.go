package db

import (
	"viandasApp/db"
	dbAddress "viandasApp/db/address"
	dbCity "viandasApp/db/city"
	dbDeliveryDriver "viandasApp/db/deliveryDriver"
	dbVehicle "viandasApp/db/vehicle"
	"viandasApp/dtos"
	"viandasApp/models"
)

func GetAllTanda() (*dtos.TandaResponse, error) {

	db := db.GetDB()

	modelTanda := []models.Tanda{}

	var allTanda dtos.TandaResponse

	if err := db.Find(&modelTanda, "active = 1").Error; err != nil {
		return nil, err
	}

	for _, valor := range modelTanda {

		deliveryDriverModel, err := dbDeliveryDriver.GetDeliveryDriverByID(valor.DeliveryDriverID)
		if err != nil {
			return nil, err
		}

		addressModel, err := dbAddress.GetAddressById(deliveryDriverModel.AddressID)
		if err != nil {
			return nil, err
		}

		vehicleModel, err := dbVehicle.GetVehicleByID(deliveryDriverModel.VehicleID)
		if err != nil {
			return nil, err
		}

		cityModel, err := dbCity.GetCityById(addressModel.CityID)
		if err != nil {
			return nil, err
		}

		tanda := dtos.TandaRes{
			ID:          valor.ID,
			Description: valor.Description,
			HourStart:   valor.HourStart,
			HourEnd:     valor.HourEnd,
			DeliveryDriver: dtos.DeliveryDriverRes{
				ID:       deliveryDriverModel.ID,
				DNI:      deliveryDriverModel.DNI,
				Name:     deliveryDriverModel.Name,
				LastName: deliveryDriverModel.LastName,
				Phone:    deliveryDriverModel.Phone,
				BornDate: deliveryDriverModel.BornDate,
				Vehicle: dtos.Vehicle{
					ID:     vehicleModel.ID,
					Brand:  vehicleModel.Brand,
					Models: vehicleModel.Models,
					Patent: vehicleModel.Patent,
					Year:   vehicleModel.Year,
				},
				Address: dtos.AddressRespone{
					ID:          addressModel.ID,
					Street:      addressModel.Street,
					Number:      addressModel.Number,
					Floor:       addressModel.Floor,
					Departament: addressModel.Departament,
					Observation: addressModel.Observation,
					Favourite:   addressModel.Favourite,
					City: dtos.AllCityResponse{
						ID:          cityModel.ID,
						Description: cityModel.Description,
						CP:          cityModel.CP,
					},
				},
			},
		}

		allTanda.Tanda = append(allTanda.Tanda, tanda)

	}

	return &allTanda, nil

}
