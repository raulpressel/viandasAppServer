package db

import (
	"viandasApp/db"
	"viandasApp/models"
)

func GetClientByIdAddress(idAddresses []int) ([]models.Client, error) {

	db := db.GetDB()

	var clientModel []models.Client

	err := db.Table("clients").
		Select("clients.id, clients.name, clients.last_name, clients.email, clients.id_user_kl, clients.phone_primary, clients.phone_secondary, clients.observation, clients.born_date").
		Joins("left JOIN client_addresses ON client_addresses.client_id = clients.id").
		Where("client_addresses.address_id IN (?)", idAddresses).
		Scan(&clientModel).Error

	return clientModel, err

}
