package db

import (
	"viandasApp/db"
	"viandasApp/models"
)

func AddAddress(addressModel models.Address, idClient int) (bool, error) {

	var clientAddressModel models.ClientAddress
	db := db.GetDB()

	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return false, err
	}

	if err := tx.Save(&addressModel).Error; err != nil {
		tx.Rollback()
		return false, err
	}

	clientAddressModel.AddressID = addressModel.ID
	clientAddressModel.ClientID = idClient

	if err := tx.Save(&clientAddressModel).Error; err != nil {
		tx.Rollback()
		return false, err
	}

	return true, tx.Commit().Error
}
