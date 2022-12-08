package db

import (
	"viandasApp/db"
	"viandasApp/models"
)

func DeleteAddress(addressModel models.Address) (bool, error) {

	db := db.GetDB()

	err := db.Save(&addressModel)

	if err.Error != nil {
		return false, err.Error
	}
	return true, err.Error

}
