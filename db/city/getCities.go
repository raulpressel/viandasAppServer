package db

import (
	"viandasApp/db"
	"viandasApp/dtos"
)

func GetCities() ([]dtos.AllCityResponse, error) {

	db := db.GetDB()

	var responseModel []dtos.AllCityResponse

	err := db.Table("cities").
		Select("cities.id, cities.description, cities.cp").
		Scan(&responseModel).Error

	return responseModel, err
}
