package db

import (
	"viandasApp/db"
	"viandasApp/models"
)

func GetFoodById(id int) (models.Food, error) {
	var db = db.ConnectDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var foodModel models.Food

	err := db.First(&foodModel, id).Error

	return foodModel, err

}
