package db

import (
	"time"
	"viandasApp/db"
	"viandasApp/dtos"
)

func GetMenuByCategory(cat int) (dtos.MenuViewer, error) {
	var db = db.ConnectDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	modelMenu := []dtos.Menu{}

	categoryMenu := []dtos.CategoryMenu{}

	foodMenu := []dtos.FoodMenu{}

	var Turns []dtos.TurnViewer

	var allMenu dtos.MenuViewer

	var dateTime time.Time = time.Now()

	err := db.Table("menus").
		Select("menus.id as menuid, turns.id as turnid, turns.description as descriptionturn  ").
		Where("? BETWEEN menus.date_start and menus.date_end", dateTime.Format("2006-01-02")).
		Joins("left JOIN turn_menus on menus.id = turn_menus.menu_id").
		Joins("left JOIN turns on turns.id = turn_menus.turn_id").
		Order("turns.id asc").
		Scan(&modelMenu).Error

	for _, valor := range modelMenu {

		err = db.Table("categories").
			Select("categories.id as category, categories.description as categorydescription, categories.title as categorytitle, categories.price as categoryprice ").
			Where("categories.active = 1 and categories.id =  ?", cat).
			Scan(&categoryMenu).Error

		turn := valor.Turnid
		menu := valor.Menuid

		var CategoryViewer []dtos.CategoryViewer

		for _, valor := range categoryMenu {

			err = db.Table("day_menus").
				Select("day_menus.date as datefood,  day_menus.food_id as foodid, foods.title as foodtitle, foods.description as fooddescription, location_imgs.location as foodurl").
				Joins("left JOIN foods ON foods.id = day_menus.food_id").
				Joins("left JOIN location_imgs on foods.location_id = location_imgs.id").
				Joins("left JOIN categories ON categories.id = foods.category_id").
				Joins("left JOIN turn_menus ON turn_menus.id = day_menus.turn_menu_id").
				Where("categories.id = ? and turn_menus.turn_id = ? and turn_menus.menu_id = ? ", valor.Category, turn, menu).
				Order("day_menus.date asc").
				Scan(&foodMenu).Error

			var Days []dtos.DayViewer

			for _, valor := range foodMenu {
				Day := dtos.DayViewer{
					Date: valor.Datefood,
					Food: dtos.FoodViewer{
						ID:          valor.Foodid,
					/* 	Title:       valor.Foodtitle,
						Description: valor.Fooddescription,
						UrlImage:    valor.Foodurl, */
					},
				}
				Days = append(Days, Day)
			}
			Category := dtos.CategoryResponse{
				ID:          valor.Category,
				Description: valor.Categorydescription,
				Title:       valor.Categorytitle,
				Price:       valor.Categoryprice,
			}

			CategoryTurn := dtos.CategoryViewer{
				Category: Category,
				Days:     Days,
			}

			CategoryViewer = append(CategoryViewer, CategoryTurn)

		}
		Turn := dtos.TurnViewer{
			ID:             valor.Turnid,
			Description:    valor.Descriptionturn,
			CategoryViewer: CategoryViewer,
		}

		Turns = append(Turns, Turn)

		allMenu = dtos.MenuViewer{
			ID:         valor.Menuid,
			TurnViewer: Turns,
		}

	}

	return allMenu, err

	/*
		err := db.Table("day_menus").
			Select("menus.id as id, menus.turn_id as turnid, turn_menus.description as descriptionturn,
			categories.id as category, categories.description as categorydescription,
			day_menus.date as datefood,  day_menus.food_id as foodid, foods.title as foodtitle, foods.description as fooddescription, location_imgs.location as foodurl").
			Joins("left JOIN foods ON foods.id = day_menus.food_id").
			Joins("left JOIN categories ON foods.category_id = categories.id").
			Joins("left JOIN menus on day_menus.menu_id = menus.id").
			Joins("left JOIN turn_menus on menus.turn_id = turn_menus.id").
			Joins("left JOIN location_imgs on foods.location_id = location_imgs.id").
			Where("? BETWEEN menus.date_start and menus.date_end", dateTime). //datetime sin horas minutos y segundos
			Scan(&modelMenu).Error */

	/* for _, valor := range modelMenu {
		responseModel = append(responseModel, *valor.ToModelResponse())
	} */

}
