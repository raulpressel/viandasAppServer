package db

import (
	"viandasApp/db"
	dbCity "viandasApp/db/city"
	"viandasApp/dtos"
	"viandasApp/models"
)

func GetAddressTakeAway() (*dtos.AddressRespone, error) {

	db := db.GetDB()

	var addressModel models.Address

	err := db.First(&addressModel, 1).Error

	cityModel, err := dbCity.GetCityById(addressModel.CityID)
	if err != nil {
		return nil, err
	}

	response := dtos.AddressRespone{
		ID:          addressModel.ID,
		Street:      addressModel.Street,
		Number:      addressModel.Number,
		Floor:       addressModel.Floor,
		Departament: addressModel.Departament,
		IDZone:      addressModel.IDZone,
		Lat:         addressModel.Lat,
		Lng:         addressModel.Lng,
		Observation: addressModel.Observation,
		City: dtos.AllCityResponse{
			ID:          cityModel.ID,
			Description: cityModel.Description,
			CP:          cityModel.CP,
		},
	}

	return &response, err

}
