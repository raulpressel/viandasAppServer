package db

import (
	"viandasApp/db"
	"viandasApp/dtos"
	"viandasApp/models"
)

func GetAllClient() (*[]dtos.ClientResponse, error) {

	db := db.GetDB()

	var pathologiesModel []models.Pathology

	var pathologiesClientModel []models.ClientPathology

	var pathology dtos.PathologyResponse

	var addressesModel []models.Address

	var address dtos.AddressRespone

	var clientResponse dtos.ClientResponse

	response := []dtos.ClientResponse{}

	var cityModel models.City

	modelClient := []models.Client{}

	err := db.Find(&modelClient).Error

	for _, client := range modelClient {

		clientResponse.Client.ID = client.ID
		clientResponse.Client.PhonePrimary = client.PhonePrimary
		clientResponse.Client.PhoneSecondary = client.PhoneSecondary
		clientResponse.Client.Name = client.Name
		clientResponse.Client.LastName = client.LastName
		clientResponse.Client.ObsClient = client.Observation
		clientResponse.Client.BornDate = client.BornDate
		clientResponse.Client.Email = client.Email

		if err := db.Table("pathologies").
			Select("pathologies.id, pathologies.description").
			//Where("client_pathologies.client_id = ?", client.ID).
			Where("pathologies.active = 1").
			Scan(&pathologiesModel).Error; err != nil {
			return nil, err
		}

		if err := db.Table("client_pathologies").
			Select("client_pathologies.id, client_pathologies.pathology_id, client_pathologies.client_id").
			//Where("client_pathologies.client_id = ?", client.ID).
			Where("client_pathologies.client_id = ?", client.ID).
			Scan(&pathologiesClientModel).Error; err != nil {
			return nil, err
		}

		for _, valor := range pathologiesModel {

			pathology.ID = valor.ID
			pathology.Description = valor.Description
			pathology.Checked = false

			for _, v := range pathologiesClientModel {
				if v.ClientID == client.ID && v.PathologyID == valor.ID {
					pathology.Checked = true
				}
			}

			clientResponse.Client.Pathologies = append(clientResponse.Client.Pathologies, pathology)

		}

		if err := db.Table("addresses").
			Select("addresses.id, addresses.street, addresses.number, addresses.floor, addresses.departament, addresses.observation, addresses.city_id, addresses.favourite").
			Joins("left JOIN client_addresses ON client_addresses.address_id = addresses.id").
			Where("client_addresses.client_id = ?", client.ID).
			Where("addresses.active = 1").
			Order("addresses.favourite desc").
			Scan(&addressesModel).Error; err != nil {
			return nil, err
		}

		for _, valor := range addressesModel {

			address.ID = valor.ID
			address.Street = valor.Street
			address.Number = valor.Number
			address.Floor = valor.Floor
			address.Departament = valor.Departament
			address.Observation = valor.Observation
			address.Favourite = valor.Favourite

			if err := db.Table("cities").
				Select("cities.id, cities.description, cities.cp ").
				Where("cities.id = ?", valor.CityID).
				Scan(&cityModel).Error; err != nil {
				return nil, err
			}

			address.City.ID = cityModel.ID
			address.City.Description = cityModel.Description
			address.City.CP = cityModel.CP

			clientResponse.Client.Address = append(clientResponse.Client.Address, address)
		}

		response = append(response, clientResponse)
	}

	return &response, err
}
