package db

import (
	"viandasApp/db"
	"viandasApp/models"
)

func GetProductById(id int) (models.Product, error) {
	db := db.GetDB()
	var model models.Product
	err := db.First(&model, id).Error
	return model, err
}
