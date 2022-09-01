package dtos

type AllFood struct {
	ID                  int
	Title               string
	Description         string
	Location            string
	Category            int
	Categorydescription string
}

type AllFoodResponse struct {
	ID          int      `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Location    string   `json:"urlImage"`
	Category    Category `json:"category"`
}

type Category struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
}

func (allFood AllFood) ToModelResponse() *AllFoodResponse {

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
