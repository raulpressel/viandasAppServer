package db

import (
	"viandasApp/db"
	"viandasApp/models"
)

func GetCityById(id int) (models.City, error) {

	db := db.GetDB()

	var cityModel models.City

	err := db.First(&cityModel, id).Error

	return cityModel, err

}
