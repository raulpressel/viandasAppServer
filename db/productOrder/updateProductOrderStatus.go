package db

import (
	"viandasApp/db"
	"viandasApp/models"
)

func GetProductOrderItemById(id int) (models.ProductOrderItem, error) {
	dbc := db.GetDB()
	var model models.ProductOrderItem
	err := dbc.First(&model, id).Error
	return model, err
}

func UpdateProductOrderItemStatus(model models.ProductOrderItem) (bool, error) {
	dbc := db.GetDB()
	err := dbc.Model(&model).Update("status", model.Status).Error
	return err == nil, err
}
