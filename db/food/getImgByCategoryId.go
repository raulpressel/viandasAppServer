package db

import (
	"viandasApp/db"
	"viandasApp/dtos"
)

func GetImgByCategoryId(id int) ([]dtos.ImgFoodByCategoryResponse, error) {

	db := db.GetDB()

	responseModel := []dtos.ImgFoodByCategoryResponse{}

	err := db.Table("food_categories").
		Select("foods.title, location_imgs.location").
		Joins("left JOIN foods ON food_categories.food_id = foods.id").
		Joins("left JOIN location_imgs ON foods.location_id = location_imgs.id").
		Where("food_categories.category_id = ?", id).
		Scan(&responseModel).Error

	return responseModel, err

}
