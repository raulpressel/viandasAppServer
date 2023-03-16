package db

import (
	"viandasApp/db"
	"viandasApp/models"
)

func DeleteTanda(tandaModel models.Tanda) (bool, error) {

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

	if err := tx.Save(&tandaModel).Error; err != nil {
		tx.Rollback()
		return false, err
	}

	if err := tx.Exec("DELETE FROM tanda_addresses WHERE tanda_id = ? ", tandaModel.ID).Error; err != nil {
		tx.Rollback()
		return false, err
	}

	return true, tx.Commit().Error

}
