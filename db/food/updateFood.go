package db

import (
	"viandasApp/db"
	"viandasApp/models"
)

func UpdateFood(foodModel models.Food, locationModel models.LocationImg, foodCategoryModel []models.FoodCategory) (bool, error) {

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

	if err := tx.Exec("DELETE FROM food_categories WHERE food_id = ?", foodModel.ID).Error; err != nil {
		tx.Rollback()
		return false, err
	}

	/* if err := tx.Save(&locationModel).Error; err != nil {
		tx.Rollback()
		return false, err
	} */

	if locationModel.Location != "" {
		if foodModel.LocationID != nil {
			if err := tx.Save(&locationModel).Error; err != nil {
				tx.Rollback()
				return false, err
			}
			foodModel.LocationID = &locationModel.ID
		} else {
			err := tx.Exec("DELETE FROM location_imgs WHERE id = ?", locationModel.ID).Error
			if err != nil {
				db.Rollback()
				return false, err
			}
		}
	}

	if err := tx.Save(&foodModel).Error; err != nil {
		tx.Rollback()
		return false, err
	}

	if err := tx.CreateInBatches(&foodCategoryModel, len(foodCategoryModel)).Error; err != nil {
		tx.Rollback()
		return false, err
	}

	return true, tx.Commit().Error

}
