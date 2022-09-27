package db

import (
	"viandasApp/db"
	"viandasApp/dtos"
)

func GetAllFood() ([]dtos.AllFoodResponse, error) {
	var db = db.ConnectDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	modelFood := []dtos.AllFood{}

	responseModelFood := []dtos.AllFoodResponse{}

	err := db.Table("foods").
		Select("foods.id, foods.title, foods.description, categories.id as category, categories.description as categorydescription, location_imgs.location").
		Joins("left JOIN location_imgs ON foods.location_id = location_imgs.id").
		Joins("left JOIN categories ON foods.category_id = categories.id").
		Where("foods.active = 1").
		Scan(&modelFood).Error

	for _, valor := range modelFood {
		responseModelFood = append(responseModelFood, *valor.ToModelResponse())
	}

	return responseModelFood, err

}