package db

import (
	"viandasApp/db"
	"viandasApp/dtos"
)

func GetAllProducts() ([]dtos.AllProductResponse, error) {
	db := db.GetDB()
	var responseModel []dtos.AllProductResponse
	err := db.Table("products").
		Select("id, title, description, price, available, product_category_id").
		Where("available = 1").
		Scan(&responseModel).Error
	return responseModel, err
}
