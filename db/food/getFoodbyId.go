package db

import (
	"viandasApp/db"
	"viandasApp/models"
)

func GetFoodById(id int) (models.Food, error) {
	/* var db = db.ConnectDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close() */

	conn := db.GetDB()

	var foodModel models.Food

	err := conn.First(&foodModel, id).Error

	return foodModel, err

}
