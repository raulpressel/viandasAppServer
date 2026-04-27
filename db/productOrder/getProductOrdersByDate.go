package db

import (
	"time"
	"viandasApp/db"
	"viandasApp/dtos"
)

func GetProductOrdersByDate(date time.Time) ([]dtos.AllProductOrderResponse, error) {
	db := db.GetDB()
	var responseModel []dtos.AllProductOrderResponse
	err := db.Table("product_orders").
		Select("id, client_name, client_last_name, product_category_title, product_title, cant, date, status").
		Where("DATE(date) = DATE(?)", date).
		Scan(&responseModel).Error
	return responseModel, err
}
