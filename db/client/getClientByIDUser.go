package db

import (
	"viandasApp/db"
	"viandasApp/dtos"
	"viandasApp/models"
)

func GetClientByIDUser(idkl string) (*dtos.ClientResponse, error) {

	db := db.GetDB()

	var pathologiesModel []models.Pathology

	var pathologiesClientModel []models.ClientPathology

	var pathology dtos.PathologyResponse

	var addressesModel []models.Address

	var address dtos.AddressRespone

	var clientResponse dtos.ClientResponse

	var cityModel models.City

	//err := db.First(&client, "id_user_kl = ?", idkl).Error

	client, res := CheckExistClient(idkl)

	if !res {

		return nil, nil
	}

	clientResponse.Client.ID = client.ID
	clientResponse.Client.PhonePrimary = client.PhonePrimary
	clientResponse.Client.PhoneSecondary = client.PhoneSecondary
	clientResponse.Client.Name = client.Name
	clientResponse.Client.LastName = client.LastName
	clientResponse.Client.ObsClient = client.Observation
	clientResponse.Client.BornDate = client.BornDate
	clientResponse.Client.Email = client.Email

	err := db.Table("pathologies").
		Select("pathologies.id, pathologies.description").
		//Where("client_pathologies.client_id = ?", client.ID).
		Where("pathologies.active = 1").
		Scan(&pathologiesModel).Error

	err = db.Table("client_pathologies").
		Select("client_pathologies.id, client_pathologies.pathology_id, client_pathologies.client_id").
		//Where("client_pathologies.client_id = ?", client.ID).
		Where("client_pathologies.client_id = ?", client.ID).
		Scan(&pathologiesClientModel).Error

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

	err = db.Table("addresses").
		Select("addresses.id, addresses.street, addresses.number, addresses.floor, addresses.departament, addresses.observation, addresses.city_id").
		Joins("left JOIN client_addresses ON client_addresses.address_id = addresses.id").
		Where("client_addresses.client_id = ?", client.ID).
		Scan(&addressesModel).Error

	for _, valor := range addressesModel {

		address.ID = valor.ID
		address.Street = valor.Street
		address.Number = valor.Number
		address.Floor = valor.Floor
		address.Departament = valor.Departament
		address.Observation = valor.Observation

		err = db.Table("cities").
			Select("cities.id, cities.description, cities.cp ").
			Where("cities.id = ?", valor.CityID).
			Scan(&cityModel).Error

		address.City.ID = cityModel.ID
		address.City.Description = cityModel.Description
		address.City.CP = cityModel.CP

		clientResponse.Client.Address = append(clientResponse.Client.Address, address)
	}

	return &clientResponse, err
}
