package db

import (
	"viandasApp/db"
	"viandasApp/models"
)

func GetProductCategoryById(id int) (models.ProductCategory, error) {
	db := db.GetDB()
	var model models.ProductCategory
	err := db.First(&model, id).Error
	return model, err
}
