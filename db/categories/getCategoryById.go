package db

import (
	"viandasApp/db"
	"viandasApp/models"
)

func GetCategoryById(id int) (models.Category, error) {
	

	db := db.GetDB()
	var categoryModel models.Category

	err := db.First(&categoryModel, id).Error

	return categoryModel, err

}
