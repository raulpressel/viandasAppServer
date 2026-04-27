package db

import (
	"viandasApp/db"
	"viandasApp/models"
)

func UploadProductOrders(orders []models.ProductOrder) ([]models.ProductOrder, error) {
	db := db.GetDB()
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return nil, err
	}
	for i := range orders {
		if err := tx.Save(&orders[i]).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	}
	return orders, tx.Commit().Error
}
