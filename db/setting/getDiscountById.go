package db

import (
	"viandasApp/db"
	"viandasApp/models"
)

func GetDiscountById(id int) (models.Discount, error) {

	db := db.GetDB()

	var discountModel models.Discount

	err := db.First(&discountModel, id).Error

	return discountModel, err

}
