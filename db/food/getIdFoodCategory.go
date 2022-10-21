package db

import (
	"viandasApp/db"
)

func GetIdFoodCategory(food int, cat int) (int, error) {
	var db = db.ConnectDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var idFoodCategory int

	err := db.Table("food_categories").
		Select("food_categories.id as menuid").
		Where("food_categories.food_id = ? and food_categories.category_id = ? ", food, cat).
		Scan(&idFoodCategory).Error

	return idFoodCategory, err

}
