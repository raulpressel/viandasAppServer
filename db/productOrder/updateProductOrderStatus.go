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
	dbc := db.GetDB()
	err := dbc.Model(&model).Update("status", model.Status).Error
	return err == nil, err
}
