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

	if err := tx.Save(&locationModel).Error; err != nil {
		tx.Rollback()
		return false, err
	}

	//categoryModel.LocationID = locationModel.ID

	if err := tx.Save(&categoryModel).Error; err != nil {
		tx.Rollback()
		return false, err
	}

	return true, tx.Commit().Error

	/* var db = db.ConnectDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	//userModel.Password, _ = EncryptPassword(userModel.Password)

	err := db.Save(&categoryModel)

	if err.Error != nil {
		return false, err.Error
	}
	return true, err.Error */



}
