package db

import (
	"viandasApp/db"
	"viandasApp/dtos"
)

func GetAllProductCategories() ([]dtos.AllProductCategoryResponse, error) {
	db := db.GetDB()
	var responseModel []dtos.AllProductCategoryResponse
	err := db.Table("product_categories").
		Select("id, title, description").
		Where("active = 1").
		Scan(&responseModel).Error
	return responseModel, err
}
