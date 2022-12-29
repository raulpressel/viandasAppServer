package db

import (
	"viandasApp/db"
	"viandasApp/models"
)

func UploadTandaAddress(tandaAddressesModel []models.TandaAddress) (bool, error) {

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

	if err := tx.CreateInBatches(&tandaAddressesModel, len(tandaAddressesModel)).Error; err != nil {
		tx.Rollback()
		return false, err
	}

	return true, tx.Commit().Error

}
