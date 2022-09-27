package dtos

type AllFood struct {
	ID                  int
	Title               string
	Description         string
	Location            string
	Category            int
	Categorydescription string
	Categorytitle       string
	Categoryprice       float32
}

type AllFoodResponse struct {
	ID          int              `json:"id"`
	Title       string           `json:"title"`
	Description string           `json:"description"`
	Location    string           `json:"urlImage"`
	Category    CategoryResponse `json:"category"`
}

type CategoryResponse struct {
	ID          int     `json:"id"`
	Description string  `json:"description"`
	Title       string  `json:"title"`
	Price       float32 `json:"price"`
}

func (allFood AllFood) ToModelResponse() *AllFoodResponse {

	modelFood := AllFoodResponse{
		ID:          allFood.ID,
		Description: allFood.Description,
		Title:       allFood.Title,
		Location:    allFood.Location,
		Category: CategoryResponse{
			ID:          allFood.Category,
			Description: allFood.Categorydescription,
			Title:       allFood.Title,
			Price:       allFood.Categoryprice,
		},
	}

	return &modelFood
}
