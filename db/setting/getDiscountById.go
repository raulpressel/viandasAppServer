package db

import (
	"sort"
	"viandasApp/db"
	"viandasApp/models"
)

func GetDiscounts() ([]models.Discount, error) {

	db := db.GetDB()

	var discountModel []models.Discount

	err := db.Find(&discountModel, "active = 1").Error

	sort.Slice(discountModel, func(i, j int) bool {
		return discountModel[i].Cant > discountModel[j].Cant
	})

	return discountModel, err

}
