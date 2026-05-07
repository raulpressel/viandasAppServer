package db

import (
	"viandasApp/db"
	"viandasApp/models"
)

func UpdateProduct(model models.Product, locationModel models.LocationImg) (bool, error) {
	conn := db.GetDB()
	tx := conn.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return false, err
	}

	if locationModel.Location != "" {
		if model.LocationID != nil {
			if err := tx.Save(&locationModel).Error; err != nil {
				tx.Rollback()
				return false, err
			}
			model.LocationID = &locationModel.ID
		} else {
			if err := tx.Exec("DELETE FROM location_imgs WHERE id = ?", locationModel.ID).Error; err != nil {
				tx.Rollback()
				return false, err
			}
		}
	}

	if err := tx.Save(&model).Error; err != nil {
		tx.Rollback()
		return false, err
	}
	return true, tx.Commit().Error
}
