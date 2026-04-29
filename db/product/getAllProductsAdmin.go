package db

import (
	"viandasApp/db"
	"viandasApp/dtos"
)

func GetAllProductsAdmin() ([]dtos.AllProductResponse, error) {
	db := db.GetDB()
	var responseModel []dtos.AllProductResponse
	err := db.Table("products").
		Select("id, title, description, price, available, product_category_id").
		Where("deleted_at IS NULL").
		Scan(&responseModel).Error
	return responseModel, err
}
