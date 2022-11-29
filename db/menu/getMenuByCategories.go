package db

import (
	"time"
	"viandasApp/db"
	"viandasApp/dtos"
)

func GetMenuByCategories(cat []int) (dtos.MenuResponse, error) {

	db := db.GetDB()

	modelMenu := []dtos.Menu{}

	foodMenuCategory := []dtos.FoodMenuCategory{}

	var turns []dtos.TurnDetailResponse

	var allMenu dtos.MenuResponse

	var dateTime time.Time = time.Now()

	err := db.Table("menus").
		Select("menus.id as menuid, menus.date_start as datestart, menus.date_end as dateend, turns.id as turnid, turns.description as descriptionturn   ").
		Where("? BETWEEN menus.date_start and menus.date_end", dateTime.Format("2006-01-02")).
		Joins("left JOIN turn_menus on menus.id = turn_menus.menu_id").
		Joins("left JOIN turns on turns.id = turn_menus.turn_id").
		Order("turns.id asc").
		Scan(&modelMenu).Error

	for _, valor := range modelMenu {

		turn := valor.Turnid
		menu := valor.Menuid

		err = db.Table("day_menus").
			Select("day_menus.id, day_menus.date as datefood,  foods.id as foodid, foods.title as foodtitle, foods.description as fooddescription, location_imgs.location as foodurl, categories.id as category, categories.description as categorydescription, categories.title as categorytitle, categories.price as categoryprice, location_imgs.location as categoryurl, categories.color as categorycolor").
			Joins("left JOIN categories ON categories.id = day_menus.category_id").
			Joins("left JOIN foods ON foods.id = day_menus.food_id").
			Joins("left JOIN location_imgs on foods.location_id = location_imgs.id").
			Joins("left JOIN turn_menus ON turn_menus.id = day_menus.turn_menu_id").
			Where("turn_menus.turn_id = ? and turn_menus.menu_id = ? ", turn, menu).
			Where("categories.active = 1 and categories.id IN (?)", cat).
			Order("day_menus.date asc").
			Scan(&foodMenuCategory).Error

		if len(foodMenuCategory) > 0 {

			var Days []dtos.DayCategoryFoodDetail

			for _, valor := range foodMenuCategory {
				Day := dtos.DayCategoryFoodDetail{
					ID:   valor.ID,
					Date: valor.Datefood,
					Food: dtos.AllFoodResponse{
						ID:          valor.Foodid,
						Title:       valor.Foodtitle,
						Description: valor.Fooddescription,
						Location:    valor.Foodurl,
						//Category: ,
					},
					Category: dtos.CategoryResponse{
						ID:          valor.Category,
						Description: valor.Categorydescription,
						Title:       valor.Categorytitle,
						Price:       valor.Categoryprice,
						Location:    valor.Categoryurl,
						Color:       valor.Categorycolor,
					},
				}
				Days = append(Days, Day)
			}

			turnDetail := dtos.TurnDetailResponse{
				ID:          valor.Turnid,
				Description: valor.Descriptionturn,
				Days:        Days,
			}

			turns = append(turns, turnDetail)

			allMenu = dtos.MenuResponse{
				Menu: dtos.MenuDetailResponse{
					ID:        valor.Menuid,
					DateStart: valor.Datestart,
					DateEnd:   valor.Dateend,
					Turn:      turns,
				},
			}

		} else {

			allMenu = dtos.MenuResponse{
				Menu: dtos.MenuDetailResponse{
					ID: 0,
				},
			}

		}
	}

	return allMenu, err

}
