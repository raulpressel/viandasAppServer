package db

import (
	"viandasApp/db"
	"viandasApp/models"
)

func GetAddressById(id int) (models.Address, error) {

	db := db.GetDB()

	var addressModel models.Address

	err := db.First(&addressModel, id).Error

	return addressModel, err

}
