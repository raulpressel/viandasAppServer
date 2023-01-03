package db

import (
	"viandasApp/db"
	"viandasApp/models"
)

func RemoveAddressToTanda(tandaAddressesModel []models.TandaAddress) (bool, error) {

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

	for _, valor := range tandaAddressesModel {
		if err := tx.Exec("DELETE FROM tanda_addresses WHERE tanda_id = ? and address_id = ?", valor.TandaID, valor.AddressID).Error; err != nil {
			tx.Rollback()
			return false, err
		}
	}

	return true, tx.Commit().Error

}
