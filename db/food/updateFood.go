package db

import (
	"viandasApp/db"
	"viandasApp/models"
)

func UpdateFood(foodModel models.Food, locationModel models.LocationImg) (bool, error) {

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

	if err := tx.Save(&foodModel).Error; err != nil {
		tx.Rollback()
		return false, err
	}

	return true, tx.Commit().Error

}
