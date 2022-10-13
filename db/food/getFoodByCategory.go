package db

import (
	"viandasApp/db"
	"viandasApp/dtos"
)

func GetFoodByCategory(cat int) ([]dtos.AllFoodResponse, error) {
	var db = db.ConnectDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	modelFood := []dtos.AllFood{}

	foodCategory := []dtos.FoodCategory{}

	responseModelFood := []dtos.AllFoodResponse{}

	categoryMenu := dtos.CategoryMenu{}

	var allFood dtos.AllFoodResponse

	if err := db.Table("foods").
		Select("foods.id, foods.title, foods.description, location_imgs.location").
		Joins("left JOIN location_imgs ON foods.location_id = location_imgs.id").
		Where("foods.active = 1").
		Scan(&modelFood).Error; err != nil {
		return nil, err
	}

	if err := db.Table("categories").
		Select("categories.id as category, categories.description as categorydescription, categories.title as categorytitle, categories.price as categoryprice ").
		Where("categories.active = 1 and categories.id = ?", cat).
		Scan(&categoryMenu).Error; err != nil {
		return nil, err
	}

	for _, valor := range modelFood {

		if err := db.Table("food_categories").
			Select("food_categories.id as category, food_categories.food_id as foodid, food_categories.category_id as categoryid ").
			Where("food_categories.food_id = ?", valor.ID).
			Scan(&foodCategory).Error; err != nil {
			return nil, err
		}

		var categoryFood dtos.CategoryResponse
		var categoriesFood []dtos.CategoryResponse

		categoryFood = dtos.CategoryResponse{
			ID:          categoryMenu.Category,
			Description: categoryMenu.Categorydescription,
			Title:       categoryMenu.Categorytitle,
			Price:       categoryMenu.Categoryprice,
			Checked:     false,
		}

		for _, fc := range foodCategory {
			if fc.Categoryid == categoryMenu.Category {
				categoryFood.Checked = true
			}
		}

		categoriesFood = append(categoriesFood, categoryFood)

		allFood = dtos.AllFoodResponse{
			ID:          valor.ID,
			Title:       valor.Title,
			Description: valor.Description,
			Location:    valor.Location,
			Category:    categoriesFood,
		}

		responseModelFood = append(responseModelFood, allFood)

	}

	return responseModelFood, nil

}
