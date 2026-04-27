package db

import (
	"viandasApp/db"
	"viandasApp/models"
)

func GetProductOrderById(id int) (models.ProductOrder, error) {
	db := db.GetDB()
	var model models.ProductOrder
	err := db.First(&model, id).Error
	return model, err
}

func UpdateProductOrderStatus(model models.ProductOrder) (bool, error) {
	db := db.GetDB()
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return false, err
	}
	if err := tx.Save(&model).Error; err != nil {
		tx.Rollback()
		return false, err
	}
	return true, tx.Commit().Error
}
