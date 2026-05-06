package db

import (
	"viandasApp/db"
)

func HasProductOrders(productTitle string) (bool, error) {
	database := db.GetDB()
	var count int64
	err := database.Table("product_order_items").
		Where("product_title = ? AND deleted_at IS NULL", productTitle).
		Count(&count).Error
	return count > 0, err
}
