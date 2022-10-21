package db

import (
	"viandasApp/db"
	"viandasApp/models"
)

func UpdateCategory(categoryModel models.Category, locationModel models.LocationImg) (bool, error) {

	var db = db.ConnectDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return false, err
	}

	if locationModel.Location != "" {
		if categoryModel.LocationID != nil {
			if err := tx.Save(&locationModel).Error; err != nil {
				tx.Rollback()
				return false, err
			}
			categoryModel.LocationID = &locationModel.ID
		} else {
			err := tx.Exec("DELETE FROM location_imgs WHERE id = ?", locationModel.ID).Error
			if err != nil {
				db.Rollback()
				return false, err
			}
		}
	}

	if err := tx.Save(&categoryModel).Error; err != nil {
		tx.Rollback()
		return false, err
	}

	return true, tx.Commit().Error

}
