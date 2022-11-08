package db

import (
	"viandasApp/db"
	"viandasApp/models"
)

func UploadCategory(categoryModel models.Category, locationModel models.LocationImg) (bool, error) {

	/* 	var db = db.ConnectDB()
	   	sqlDB, _ := db.DB()
	   	defer sqlDB.Close() */

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

	if locationModel.Location != "" {
		if err := tx.Save(&locationModel).Error; err != nil {
			tx.Rollback()
			return false, err
		}
		categoryModel.LocationID = &locationModel.ID

	}

	if err := tx.Save(&categoryModel).Error; err != nil {
		tx.Rollback()
		return false, err
	}

	return true, tx.Commit().Error

}
