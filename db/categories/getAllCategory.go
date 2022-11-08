package db

import (
	"viandasApp/db"
	"viandasApp/dtos"
)

func GetAllCategory() ([]dtos.AllCategoryResponse, error) {

	db := db.GetDB()

	var responseModel []dtos.AllCategoryResponse

	err := db.Table("categories").
		Select("categories.id, categories.description, categories.title, categories.color, categories.price, location_imgs.location").
		Joins("left JOIN location_imgs ON categories.location_id = location_imgs.id").
		Where("categories.active = 1").
		Scan(&responseModel).Error

	return responseModel, err

}
