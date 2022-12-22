package db

import (
	"viandasApp/db"
	dbAddress "viandasApp/db/address"
	dbCity "viandasApp/db/city"
	dbVehicle "viandasApp/db/vehicle"
	"viandasApp/dtos"
	"viandasApp/models"
)

func GetAllDeliveryDriver() (dtos.DeliveryDriverResponse, error) {

	db := db.GetDB()

	modelDeliveryDriver := []models.DeliveryDriver{}

	var allDeliveryDriver dtos.DeliveryDriverResponse

	if err := db.Find(&modelDeliveryDriver, "active = 1").Error; err != nil {
		return allDeliveryDriver, err
	}

	for _, valor := range modelDeliveryDriver {

		addressModel, err := dbAddress.GetAddressById(valor.AddressID)
		if err != nil {
			return allDeliveryDriver, err
		}

		vehicleModel, err := dbVehicle.GetVehicleByID(valor.VehicleID)
		if err != nil {
			return allDeliveryDriver, err
		}

		cityModel, err := dbCity.GetFoodById(addressModel.CityID)
		if err != nil {
			return allDeliveryDriver, err
		}

		deliveryDriver := dtos.DeliveryDriverRes{
			ID:       valor.DNI,
			DNI:      valor.DNI,
			Name:     valor.Name,
			LastName: valor.LastName,
			Phone:    valor.Phone,
			BornDate: valor.BornDate,
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
		}

		allDeliveryDriver.DeliveryDriver = append(allDeliveryDriver.DeliveryDriver, deliveryDriver)

	}

	return allDeliveryDriver, nil

}
