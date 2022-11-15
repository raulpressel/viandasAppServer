package db

import (
	"time"
	"viandasApp/db"
	"viandasApp/dtos"
)

func GetMenuActive() (dtos.MenuViewer, error) {
	/* var db = db.ConnectDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close() */

	db := db.GetDB()

	modelMenu := []dtos.Menu{}

	categoryMenu := []dtos.CategoryMenu{}

	var Turns []dtos.TurnViewer

	var allMenu dtos.MenuViewer

	var dateTime time.Time = time.Now()

	err := db.Table("menus").
		Select("menus.id as menuid, menus.date_start as datestart, menus.date_end as dateend, turns.id as turnid, turns.description as descriptionturn   ").
		Where("? BETWEEN menus.date_start and menus.date_end", dateTime.Format("2006-01-02")).
		Joins("left JOIN turn_menus on menus.id = turn_menus.menu_id").
		Joins("left JOIN turns on turns.id = turn_menus.turn_id").
		Order("turns.id asc").
		Scan(&modelMenu).Error

	/* 	subQuery1 := db.Model(&models.Menu{}).Where("? BETWEEN menus.date_start and menus.date_end", dateTime.Format("2006-01-02"))
	   	fmt.Println(subQuery1) */

	for _, valor := range modelMenu {

		err = db.Table("categories").
			Select("categories.id as category, categories.description as categorydescription, categories.title as categorytitle, categories.price as categoryprice, location_imgs.location as categoryurl, categories.color as categorycolor").
			Joins("left JOIN location_imgs on categories.location_id = location_imgs.id").
			Where("categories.active = 1").
			Scan(&categoryMenu).Error

		turn := valor.Turnid
		menu := valor.Menuid

		var CategoryViewer []dtos.CategoryViewer

		for _, valor := range categoryMenu {

			foodMenu := []dtos.FoodMenu{}

			err = db.Table("day_menus").
				Select("day_menus.date as datefood,  foods.id as foodid, foods.title as foodtitle, foods.description as fooddescription, location_imgs.location as foodurl").
				Joins("left JOIN categories ON categories.id = day_menus.category_id").
				Joins("left JOIN foods ON foods.id = day_menus.food_id").
				Joins("left JOIN location_imgs on foods.location_id = location_imgs.id").
				Joins("left JOIN turn_menus ON turn_menus.id = day_menus.turn_menu_id").
				Where("day_menus.category_id = ? and turn_menus.turn_id = ? and turn_menus.menu_id = ? ", valor.Category, turn, menu).
				Order("day_menus.date asc").
				Scan(&foodMenu).Error

			var Days []dtos.DayViewer

			if len(foodMenu) > 0 {

				for _, valor := range foodMenu {
					Day := dtos.DayViewer{
						Date: valor.Datefood,
						Food: dtos.FoodViewer{
							ID:          valor.Foodid,
							Title:       valor.Foodtitle,
							Description: valor.Fooddescription,
							UrlImage:    valor.Foodurl,
						},
					}
					Days = append(Days, Day)
				}
				Category := dtos.CategoryResponse{
					ID:          valor.Category,
					Description: valor.Categorydescription,
					Title:       valor.Categorytitle,
					Location:    valor.Categoryurl,
					Price:       valor.Categoryprice,
					Color:       valor.Categorycolor,
				}

				CategoryTurn := dtos.CategoryViewer{
					Category: Category,
					Days:     Days,
				}

				CategoryViewer = append(CategoryViewer, CategoryTurn)
			}

		}
		Turn := dtos.TurnViewer{
			ID:             valor.Turnid,
			Description:    valor.Descriptionturn,
			CategoryViewer: CategoryViewer,
		}

		Turns = append(Turns, Turn)

		allMenu = dtos.MenuViewer{
			ID:         valor.Menuid,
			DateStart:  valor.Datestart,
			DateEnd:    valor.Dateend,
			TurnViewer: Turns,
		}

	}

	return allMenu, err

}
