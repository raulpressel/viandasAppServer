package dtos

type AllCategoryResponse struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Title       string `json:"title"`
	Price       string `json:"price"`
}
