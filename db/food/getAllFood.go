package db

import (
	"viandasApp/db"
	"viandasApp/dtos"
)

func GetAllFood() ([]dtos.AllFoodResponse, error) {
	var db = db.ConnectDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	modelFood := []dtos.AllFood{}

	foodCategory := []dtos.FoodCategory{}

	responseModelFood := []dtos.AllFoodResponse{}

	categoryMenu := []dtos.CategoryMenu{}

	var allFood dtos.AllFoodResponse

	err := db.Table("foods").
		Select("foods.id, foods.title, foods.description, location_imgs.location").
		Joins("left JOIN location_imgs ON foods.location_id = location_imgs.id").
		//Joins("left JOIN food_categories ON food_categories.food_id = foods.id").
		//Joins("left JOIN categories ON categories.id = food_categories.category_id").
		Where("foods.active = 1").
		Scan(&modelFood).Error

	err = db.Table("categories").
		Select("categories.id as category, categories.description as categorydescription, categories.title as categorytitle, categories.price as categoryprice ").
		Where("categories.active = 1").
		Scan(&categoryMenu).Error

	for _, valor := range modelFood {

		err = db.Table("food_categories").
			Select("food_categories.id as category, food_categories.food_id as foodid, food_categories.category_id as categoryid ").
			Where("food_categories.food_id = ?", valor.ID).
			Scan(&foodCategory).Error

		var categoryFood dtos.CategoryResponse
		var categoriesFood []dtos.CategoryResponse

		//cantFoodCategory := len(foodCategory)

		for _, cat := range categoryMenu {

			//for _, fc := range foodCategory {
			categoryFood = dtos.CategoryResponse{
				ID:          cat.Category,
				Description: cat.Categorydescription,
				Title:       cat.Categorytitle,
				Price:       cat.Categoryprice,
				Checked:     false,
			}

			for _, fc := range foodCategory {
				if fc.Categoryid == cat.Category {
					categoryFood.Checked = true
				}
			}

			//}

			categoriesFood = append(categoriesFood, categoryFood)

		}

		allFood = dtos.AllFoodResponse{
			ID:          valor.ID,
			Title:       valor.Title,
			Description: valor.Description,
			Location:    valor.Location,
			Category:    categoriesFood,
		}

		responseModelFood = append(responseModelFood, allFood)

	}

	/* for _, valor := range modelFood {
		responseModelFood = append(responseModelFood, *valor.ToModelResponse())
	} */

	return responseModelFood, err

}
