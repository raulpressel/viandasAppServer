package dtos

import "time"

type AllMenu struct {
	ID              int
	Turnid          int
	Descriptionturn string
}
type CategoryMenu struct {
	Category            int
	Categorydescription string
	Categorytitle       string
	Categoryprice       float32
}

type FoodMenu struct {
	Datefood        time.Time
	Foodid          int
	Foodtitle       string
	Fooddescription string
	Foodurl         string
}

/* type MonthMenuRequest struct {
	ID     int               `json:"id"`
	Active bool              `json:"active"`
	Week   []WeekMenuRequest `json:"weeks"`
}

type WeekMenuRequest struct {
	Name      string        `json:"name"`
	DateStart string        `json:"dateStart"`
	DateEnd   string        `json:"dateEnd"`
	Menu      []MenuRequest `json:"menu"`
} */

type TurnMenuRequest struct {
	Menu []MenuRequest `json:"turns"`
}

type MenuRequest struct {
	TurnId    int              `json:"id"`
	DateStart string           `json:"dateStart"`
	DateEnd   string           `json:"dateEnd"`
	DayMenu   []DayMenuRequest `json:"days"`
}

type DayMenuRequest struct {
	Date string `json:"date"`
	Food int    `json:"idFood"`
}

//////////////////////////////RESPONSE

type MenuViewer struct {
	ID         int          `json:"id"`
	TurnViewer []TurnViewer `json:"turnsViewer"`
}

type TurnViewer struct {
	ID             int              `json:"id"`
	Description    string           `json:"description"`
	CategoryViewer []CategoryViewer `json:"categoryViewer"`
}

type CategoryViewer struct {
	Category CategoryResponse `json:"category"`
	Days     []DayViewer      `json:"daysViewer"`
}

type DayViewer struct {
	Date time.Time  `json:"date"`
	Food FoodViewer `json:"foodViewer"`
}

type FoodViewer struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	UrlImage    string `json:"urlImage"`
}

func (allMenu AllMenu) ToModelResponse() *MenuViewer {

	/* 	Food := FoodResponse{
		ID:          allMenu.Foodid,
		Title:       allMenu.Foodtitle,
		Description: allMenu.Fooddescription,
		UrlImage:    allMenu.Foodurl,
	} */

	/* 	Day := DayResponse{
	   		Date: allMenu.Datefood,
	   		Food: FoodResponse{
	   			ID:          allMenu.Foodid,
	   			Title:       allMenu.Foodtitle,
	   			Description: allMenu.Fooddescription,
	   			UrlImage:    allMenu.Foodurl,
	   		},
	   	}

	   	Category := CategoryResponse{
	   		ID:          allMenu.Category,
	   		Description: allMenu.Categorydescription,
	   	}
		var Days []DayResponse

	   	Days = append(Days, Day)

	   	CategoryTurn := CategoryTurnResponse{
	   		Category,
	   		Days,
	   	}
	*/
	/*

		var CategoryTurns []CategoryTurnResponse

		CategoryTurns = append(CategoryTurns, CategoryTurn)

		Turn := TurnResponse{
			allMenu.Turnid,
			allMenu.Descriptionturn,
			CategoryTurns,
		}

		var Turns []TurnResponse

		Turns = append(Turns, Turn)

		modelMenu := AllMenuResponse{
			ID:          allMenu.ID,
			TurnRespone: Turns,
		}

		return &modelMenu */

	return nil
}

/*
modelMenu := AllMenuResponse{
	ID: allMenu.ID,
	TurnRespone: TurnResponse{
		ID:          allMenu.Turnid,
		Description: allMenu.Descriptionturn,
		CategoryTurn: CategoryTurn{
			Category: Category,
			Days: Day{
				Date: allMenu.Datefood,
				//Food: FoodsAux,

				Food: FoodResponse{
					ID:          allMenu.Foodid,
					Title:       allMenu.Foodtitle,
					Description: allMenu.Fooddescription,
					UrlImage:    allMenu.Foodurl,
				},
			},
		},
	}, */
