package dtos

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

type MenuRequest struct {
	TurnId    int              `json:"turnId"`
	DateStart string           `json:"dateStart"`
	DateEnd   string           `json:"dateEnd"`
	DayMenu   []DayMenuRequest `json:"days"`
}

type DayMenuRequest struct {
	Date string `json:"date"`
	Food int    `json:"foodId"`
}

type MenuResponse struct {
	Date                string `json:"date"`
	TurnId              int    `json:"turnId"`
	Food                int    `json:"foodId"`
	Foodtitle           string `json:"foodTitle"`
	Categoryid          int    `json:"categoryId"`
	Categorydescription string `json:"categoryDescription"`
}

/* func (allFood AllFood) ToModelResponse() *AllFoodResponse {

	modelFood := AllFoodResponse{
		ID:          allFood.ID,
		Description: allFood.Description,
		Title:       allFood.Title,
		Location:    allFood.Location,
		Category: Category{
			ID:          allFood.Category,
			Description: allFood.Categorydescription,
		},
	}

	return &modelFood
}
*/
