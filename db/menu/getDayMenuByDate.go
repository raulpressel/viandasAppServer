package db

import (
	"time"
	"viandasApp/db"
	"viandasApp/dtos"
)

func GetDayMenuByDate(date time.Time) ([]dtos.DayMenuResponse, error) {

	db := db.GetDB()

	dayMenuDto := []dtos.DayMenuDateDto{}

	dayMenuResponse := []dtos.DayMenuResponse{}

	var categoryMenu dtos.CategoryMenu

	err := db.Table("day_menus").
		Select("day_menus.id as id, day_menus.date as date, foods.id as foodid, foods.title as foodtitle, foods.description as fooddescription , categories.id as categoryid, categories.description as categorydescription, categories.title as categorytitle , categories.price as categoryprice, location_imgs.location as foodlocation").
		Joins("left JOIN turn_menus on turn_menus.id = day_menus.turn_menu_id").
		Joins("left JOIN menus on menus.id = turn_menus.menu_id").
		Joins("left JOIN categories ON categories.id = day_menus.category_id").
		Joins("left JOIN foods ON foods.id = day_menus.food_id").
		Joins("left JOIN location_imgs ON foods.location_id = location_imgs.id").
		Where("day_menus.date = ?", date.Format("2006-01-02")).
		Scan(&dayMenuDto).Error

	for _, valor := range dayMenuDto {

		var categoriesFood []dtos.CategoryResponse

		for _, cat := range dayMenuDto {
			err = db.Table("categories").
				Select("categories.id as category, categories.description as categorydescription, categories.title as categorytitle, categories.price as categoryprice ").
				Where("categories.id = ? ", cat.Categoryid).
				Scan(&categoryMenu).Error

			categoryFood := dtos.CategoryResponse{
				ID:          categoryMenu.Category,
				Description: categoryMenu.Categorydescription,
				Title:       categoryMenu.Categorytitle,
				Price:       categoryMenu.Categoryprice,
				Checked:     true,
			}

			categoriesFood = append(categoriesFood, categoryFood)
		}

		dayFoodMenu := dtos.DayFoodMenuResponse{
			ID:          valor.Foodid,
			Title:       valor.Foodtitle,
			Description: valor.Fooddescription,
			Location:    valor.Foodlocation,
			Categories:  categoriesFood,
		}

		dayCategoryMenu := dtos.CategoryResponse{
			ID:          valor.Categoryid,
			Description: valor.Categorydescription,
			Title:       valor.Categorytitle,
			Price:       valor.Categoryprice,
		}

		dayMenu := dtos.DayMenuResponse{
			ID:       valor.ID,
			Date:     valor.Date,
			Food:     dayFoodMenu,
			Category: dayCategoryMenu,
		}

		dayMenuResponse = append(dayMenuResponse, dayMenu)

	}

	return dayMenuResponse, err

}
