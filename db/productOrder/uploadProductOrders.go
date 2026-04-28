package db

import (
	"viandasApp/db"
	"viandasApp/models"
)

func UploadProductOrder(order models.ProductOrder) (models.ProductOrder, error) {
	dbc := db.GetDB()
	tx := dbc.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return models.ProductOrder{}, err
	}
	if err := tx.Create(&order).Error; err != nil {
		tx.Rollback()
		return models.ProductOrder{}, err
	}
	return order, tx.Commit().Error
}
