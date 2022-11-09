package db

import (
	"viandasApp/db"
	"viandasApp/models"
)

func GetCategoryById(id int) (models.Pathology, error) {
	db := db.GetDB()

	var categoryPathology models.Pathology

	err := db.First(&categoryPathology, id).Error

	return categoryPathology, err

}
