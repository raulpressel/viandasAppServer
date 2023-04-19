package db

import (
	"viandasApp/db"
	"viandasApp/dtos"
	"viandasApp/models"
)

func GetAllDiscount() (*dtos.DiscountResponse, error) {

	db := db.GetDB()

	discountModel := []models.Discount{}

	var allDiscount dtos.DiscountResponse

	if err := db.Find(&discountModel, "active = 1").Error; err != nil {
		return nil, err
	}

	for _, valor := range discountModel {

		discounts := dtos.Discount{
			ID:          valor.ID,
			Description: valor.Description,
			Cant:        valor.Cant,
			Percentage:  valor.Percentage,
		}

		allDiscount.Discount = append(allDiscount.Discount, discounts)

	}

	return &allDiscount, nil

}
