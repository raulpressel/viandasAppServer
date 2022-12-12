package dtos

type AllFood struct {
	ID          int
	Title       string
	Description string
	Location    string
	Category    int
}

type FoodCategory struct {
	ID         int
	Foodid     int
	Categoryid int
}

type FoodResponse struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Location    string `json:"urlImage"`
}

type AllFoodResponse struct {
	ID          int                `json:"id"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	Location    string             `json:"urlImage"`
	Category    []CategoryResponse `json:"categories"`
}

type CategoryResponse struct {
	ID          int     `json:"id"`
	Description string  `json:"description"`
	Title       string  `json:"title"`
	Price       float32 `json:"price"`
	Location    string  `json:"urlImage"`
	Checked     bool    `json:"checked"`
	Color       string  `json:"color"`
}

type ImgFoodByCategoryResponse struct {
	Title    string `json:"title"`
	Location string `json:"urlImage"`
}

/* func (allFood AllFood) ToModelResponse() *AllFoodResponse {

	modelFood := AllFoodResponse{
		ID:          allFood.ID,
		Description: allFood.Description,
		Title:       allFood.Title,
		Location:    allFood.Location,
		Category: CategoryResponse{
			ID:          allFood.Category,
			Description: allFood.Categorydescription,
			Title:       allFood.Categorytitle,
			Price:       allFood.Categoryprice,
		},
	}

	return &modelFood
} */
