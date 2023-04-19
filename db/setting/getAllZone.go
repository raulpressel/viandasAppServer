package db

import (
	"viandasApp/db"
	"viandasApp/dtos"
	"viandasApp/models"
)

func GetAllZone() (*dtos.ZoneResponse, error) {

	db := db.GetDB()

	zoneModel := []models.Zone{}

	var allZones dtos.ZoneResponse

	if err := db.Find(&zoneModel, "active = 1").Error; err != nil {
		return nil, err
	}

	for _, valor := range zoneModel {

		zones := dtos.Zone{
			ID:          valor.ID,
			Description: valor.Description,
			Price:       valor.Price,
		}

		allZones.Zone = append(allZones.Zone, zones)

	}

	return &allZones, nil

}
