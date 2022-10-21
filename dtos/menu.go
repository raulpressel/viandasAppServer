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
	Foodcategory        int
}
type Menu struct {
	Menuid          int
	Turnid          int
	Datestart       time.Time
	Dateend         time.Time
	Descriptionturn string
}
type CategoryMenu struct {
	Category            int
	Categorydescription string
	Categorytitle       string
	Categoryprice       float32
	Categoryurl         string
}

type FoodMenu struct {
	Datefood        time.Time
	Foodid          int
	Foodtitle       string
	Fooddescription string
	Foodurl         string
}

type ValidateDateMenuRequest struct {
	DateStart string `json:"dateStart"`
	DateEnd   string `json:"dateEnd"`
}

type ValidateDateMenuRespone struct {
	ValidDateMenu bool `json:"validDateMenu"`
}
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
	Date     string `json:"date"`
	Food     int    `json:"idFood"`
	Category int    `json:"idCategory"`
}

type DayDateMenuRequest struct {
	Date string `json:"date"`
}

//////////////////////////////RESPONSE

type MenuViewer struct {
	ID         int          `json:"id"`
	DateStart  time.Time    `json:"dateStart"`
	DateEnd    time.Time    `json:"dateEnd"`
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
	IdDayMenu  int `json:"idDay"`
	IdFood     int `json:"idFood"`
	IdCategory int `json:"idCategory"`
}

type DayMenuResponse struct {
	ID       int                 `json:"id"`
	Date     time.Time           `json:"date"`
	Food     DayFoodMenuResponse `json:"food"`
	Category CategoryResponse    `json:"category"`
}
type DayFoodMenuResponse struct {
	ID          int                `json:"id"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	Location    string             `json:"urlImage"`
	Categories  []CategoryResponse `json:"categories"`
}

type DayCategoryMenuResponse struct {
	ID          int     `json:"id"`
	Description string  `json:"description"`
	Title       string  `json:"title"`
	Price       float32 `json:"price"`
}

type AllMenu struct {
	Menuid          int
	Menudatestart   time.Time
	Menudateend     time.Time
	Turnid          int
	IsCurrent       bool
	Turndescription string
}

type AllMenuResponse struct {
	ID        int              `json:"menuId"`
	DateStart time.Time        `json:"dateStart"`
	DateEnd   time.Time        `json:"dateEnd"`
	IsCurrent bool             `json:"isCurrent"`
	Turn      TurnMenuResponse `json:"turn"`
}

type TurnMenuResponse struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
}

func (allMenu AllMenu) ToAllMenuResponse() *AllMenuResponse {

	allMenuResponse := AllMenuResponse{
		ID:        allMenu.Menuid,
		DateStart: allMenu.Menudatestart,
		DateEnd:   allMenu.Menudateend,
		IsCurrent: allMenu.IsCurrent,
		Turn: TurnMenuResponse{
			ID:          allMenu.Turnid,
			Description: allMenu.Turndescription,
		},
	}

	return &allMenuResponse
}
