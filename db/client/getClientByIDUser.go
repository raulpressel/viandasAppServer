package db

import (
	"viandasApp/db"
	"viandasApp/dtos"
	"viandasApp/models"
)

func GetClientByIDUser(idkl string) (dtos.ClientRespone, error) {

	db := db.GetDB()

	var pathologiesModel []models.Pathology

	var pathology dtos.PathologyResponse

	var addressesModel []models.Address

	var address dtos.AddressRespone

	var clientResponse dtos.ClientRespone

	//err := db.First(&client, "id_user_kl = ?", idkl).Error

	client, res := CheckExistClient(idkl)

	if !res {

		return clientResponse, nil
	}

	clientResponse.ID = client.ID
	clientResponse.PhonePrimary = client.PhonePrimary
	clientResponse.PhoneSecondary = client.PhoneSecondary
	clientResponse.Name = client.Name
	clientResponse.LastName = client.LastName
	clientResponse.ObsClient = client.Observation
	clientResponse.BornDate = client.BornDate

	err := db.Table("pathologies").
		Select("pathologies.id, pathologies.description").
		Joins("left JOIN client_pathologies ON client_pathologies.pathology_id = pathologies.id").
		Where("client_pathologies.client_id = ?", client.ID).
		Scan(&pathologiesModel).Error

	for _, valor := range pathologiesModel {

		pathology.ID = valor.ID
		pathology.Description = valor.Description
		pathology.Checked = true
		clientResponse.Pathologies = append(clientResponse.Pathologies, pathology)

	}

	err = db.Table("addresses").
		Select("addresses.id, addresses.street, addresses.number, addresses.floor, addresses.departament, addresses.observation").
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

		clientResponse.Address = append(clientResponse.Address, address)
	}

	return clientResponse, err
}
