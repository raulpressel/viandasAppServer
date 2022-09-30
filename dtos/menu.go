package dtos

import (
	"time"
)

type DayMenuDateDto struct {
	ID                  int
	Date                time.Time
	Foodid              int
	Foodtitle           string
	Fooddescription     string
	Foodlocation        string
	Categoryid          int
	Categorydescription string
	Categorytitle       string
	Categoryprice       float32
}
type AllMenu struct {
	Menuid          int
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

type DayDateMenuRequest struct {
	Date string `json:"date"`
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

type DayMenuEditRequest struct {
	IdDayMenu int `json:"idDay"`
	IdFood    int `json:"idFood"`
}

type DayMenuResponse struct {
	ID   int                 `json:"id"`
	Date time.Time           `json:"date"`
	Food DayFoodMenuResponse `json:"food"`
}
type DayFoodMenuResponse struct {
	ID          int                     `json:"id"`
	Title       string                  `json:"title"`
	Description string                  `json:"description"`
	Location    string                  `json:"urlImage"`
	Category    DayCategoryMenuResponse `json:"category"`
}

type DayCategoryMenuResponse struct {
	ID          int     `json:"id"`
	Description string  `json:"description"`
	Title       string  `json:"title"`
	Price       float32 `json:"price"`
}

func (dayMenuDateDto DayMenuDateDto) ToDayMenuDateResponse() *DayMenuResponse {

	dayMenuResponse := DayMenuResponse{
		ID:   dayMenuDateDto.ID,
		Date: dayMenuDateDto.Date,
		Food: DayFoodMenuResponse{
			ID:          dayMenuDateDto.Foodid,
			Title:       dayMenuDateDto.Foodtitle,
			Description: dayMenuDateDto.Fooddescription,
			Location:    dayMenuDateDto.Foodlocation,
			Category: DayCategoryMenuResponse{
				ID:          dayMenuDateDto.Categoryid,
				Description: dayMenuDateDto.Categorydescription,
				Title:       dayMenuDateDto.Categorytitle,
				Price:       dayMenuDateDto.Categoryprice,
			},
		},
	}

	return &dayMenuResponse
}
