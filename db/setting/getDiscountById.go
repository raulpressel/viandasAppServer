package db

import (
	"viandasApp/db"
	"viandasApp/models"
)

func GetDiscounts() ([]models.Discount, error) {

	db := db.GetDB()

	var discountModel []models.Discount

	err := db.Find(&discountModel, "active = 1").Error

	return discountModel, err

}
