package db

import (
	"viandasApp/db"
	"viandasApp/models"
)

func RegisterClient(clientModel models.Client, clientPathologyModel []models.ClientPathology, addressModel []models.Address) (bool, error) {

	var clientAddrModel []models.ClientAddress
	var cAModel models.ClientAddress

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

	if err := tx.Save(&clientModel).Error; err != nil {
		tx.Rollback()
		return false, err
	}

	if len(clientPathologyModel) > 0 {

		for i := range clientPathologyModel {

			clientPathologyModel[i].ClientID = clientModel.ID
		}

		if err := tx.CreateInBatches(&clientPathologyModel, len(clientPathologyModel)).Error; err != nil {
			tx.Rollback()
			return false, err
		}
	}

	if err := tx.CreateInBatches(&addressModel, len(addressModel)).Error; err != nil {
		tx.Rollback()
		return false, err
	}

	for _, val := range addressModel {

		cAModel.ClientID = clientModel.ID
		cAModel.AddressID = val.ID

		clientAddrModel = append(clientAddrModel, cAModel)

	}

	if err := tx.CreateInBatches(&clientAddrModel, len(clientAddrModel)).Error; err != nil {
		tx.Rollback()
		return false, err
	}

	return true, tx.Commit().Error
}
