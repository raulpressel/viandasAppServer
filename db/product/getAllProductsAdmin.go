package db

import (
	"viandasApp/db"
	"viandasApp/dtos"
)

func GetAllProductsAdmin() ([]dtos.AllProductResponse, error) {
	conn := db.GetDB()
	var responseModel []dtos.AllProductResponse
	err := conn.Table("products").
		Select("products.id, products.title, products.description, products.price, products.available, products.product_category_id, COALESCE(location_imgs.location, '') as url_image").
		Joins("LEFT JOIN location_imgs ON location_imgs.id = products.location_id").
		Where("products.deleted_at IS NULL").
		Scan(&responseModel).Error
	return responseModel, err
}
