package db

import (
	"viandasApp/db"
	"viandasApp/dtos"
	"viandasApp/models"
)

func GetAllClientByTandas(tandas []int) (*[]dtos.Client, error) {

	db := db.GetDB()

	var pathologiesModel []models.Pathology

	var pathologiesClientModel []models.ClientPathology

	var address dtos.AddressRespone

	var pathology dtos.PathologyResponse

	response := []dtos.Client{}

	var cityModel models.City

	/* var modelTandaAddress []models.TandaAddress

	err := db.Find(&modelTandaAddress, "tanda_id IN (?)", tandas).Error

	var idAddresses []int

	for _, val := range modelTandaAddress {
		idAddresses = append(idAddresses, val.AddressID)
	}

	modelClient, err := GetClientByIdAddress(idAddresses) */

	var modelClient []models.Client

	err := db.Table("clients").
		Select("distinct clients.id, clients.name, clients.last_name, clients.email, clients.id_user_kl, clients.phone_primary, clients.phone_secondary, clients.observation, clients.born_date").
		Joins("JOIN client_addresses ON client_addresses.client_id = clients.id").
		Where("client_addresses.address_id IN (select address_id from tanda_addresses where tanda_id IN (?))", tandas).
		Scan(&modelClient).Error

	for _, client := range modelClient {

		var clientResponse dtos.Client

		clientResponse.ID = client.ID
		clientResponse.PhonePrimary = client.PhonePrimary
		clientResponse.PhoneSecondary = client.PhoneSecondary
		clientResponse.Name = client.Name
		clientResponse.LastName = client.LastName
		clientResponse.ObsClient = client.Observation
		clientResponse.BornDate = client.BornDate
		clientResponse.Email = client.Email

		notesClientModel, _ := GetNoteByClientId(client.ID)

		clientResponse.Note.ID = notesClientModel.ID
		clientResponse.Note.Note = notesClientModel.Note

		if err := db.Table("pathologies").
			Select("pathologies.id, pathologies.description").
			Where("pathologies.active = 1").
			Scan(&pathologiesModel).Error; err != nil {
			return nil, err
		}

		if err := db.Table("client_pathologies").
			Select("client_pathologies.id, client_pathologies.pathology_id, client_pathologies.client_id").
			Where("client_pathologies.client_id = ?", client.ID).
			Scan(&pathologiesClientModel).Error; err != nil {
			return nil, err
		}

		var pathologies []dtos.PathologyResponse

		for _, valor := range pathologiesModel {

			pathology.ID = valor.ID
			pathology.Description = valor.Description
			pathology.Checked = false

			for _, v := range pathologiesClientModel {
				if v.ClientID == client.ID && v.PathologyID == valor.ID {
					pathology.Checked = true
				}
			}

			pathologies = append(pathologies, pathology)

		}

		clientResponse.Pathologies = pathologies

		var addressesModel []models.Address

		if err := db.Table("addresses").
			Select("addresses.id, addresses.street, addresses.number, addresses.floor, addresses.departament, addresses.observation, addresses.city_id, addresses.favourite").
			Joins("left JOIN client_addresses ON client_addresses.address_id = addresses.id").
			Where("client_addresses.client_id = ?", client.ID).
			Where("addresses.id IN (select address_id from tanda_addresses where tanda_id IN (?))", tandas).
			Where("addresses.active = 1").
			Order("addresses.favourite desc").
			Scan(&addressesModel).Error; err != nil {
			return nil, err
		}

		var addresses []dtos.AddressRespone

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

			addresses = append(addresses, address)
		}

		clientResponse.Address = addresses

		response = append(response, clientResponse)
	}

	return &response, err
}
