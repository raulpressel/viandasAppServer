package db

import (
	"viandasApp/db"
	"viandasApp/models"
)

func DeleteDiscount(discountModel models.Discount) (bool, error) {

	db := db.GetDB()

	err := db.Save(&discountModel)

	if err.Error != nil {
		return false, err.Error
	}
	return true, err.Error

}
