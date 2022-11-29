package db

import (
	"viandasApp/db"
	"viandasApp/models"
)

func UpdateClient(clientModel models.Client, clientPathologyModel []models.ClientPathology) (bool, error) {

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

	if err := tx.Exec("DELETE FROM client_pathologies WHERE client_id = ?", clientModel.ID).Error; err != nil {
		tx.Rollback()
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

	return true, tx.Commit().Error
}
