package db

import (
	"viandasApp/db"
	"viandasApp/models"
)

func UpdateAddressesZone(addressesModel []models.Address) (bool, error) {

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

	for _, address := range addressesModel {
		if err := tx.Model(&models.Address{}).Where("id = ?", address.ID).Update("id_zone", address.IDZone).Error; err != nil {
			tx.Rollback()
			return false, err
		}
	}

	return true, tx.Commit().Error
}
