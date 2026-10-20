package db

import (
	"viandasApp/db"
	"viandasApp/models"
)

func GetAllActiveAddresses() ([]models.Address, error) {

	db := db.GetDB()

	var addressesModel []models.Address

	err := db.Where("active = 1").Where("id <> ?", 100).Find(&addressesModel).Error

	return addressesModel, err

}
