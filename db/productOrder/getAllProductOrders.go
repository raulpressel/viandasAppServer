package db

import (
	"viandasApp/db"
	"viandasApp/dtos"
)

func GetAllProductOrders() ([]dtos.AllProductOrderResponse, error) {
	db := db.GetDB()
	var responseModel []dtos.AllProductOrderResponse
	err := db.Table("product_orders").
		Select("id, client_name, client_last_name, product_category_title, product_title, cant, date, status").
		Scan(&responseModel).Error
	return responseModel, err
}
