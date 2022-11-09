package db

import (
	"viandasApp/db"
	"viandasApp/models"
)

func GetAllPathology() ([]models.Pathology, error) {

	db := db.GetDB()

	var responseModel []models.Pathology

	/* err := db.Table("pathologies").
	Select("categories.id, categories.description, categories.title, categories.price").
	Where("categories.active = 1").
	Scan(&responseModel).Error */

	err := db.Find(&responseModel).Error

	return responseModel, err

}
