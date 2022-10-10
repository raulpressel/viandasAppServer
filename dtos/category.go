package dtos

import "viandasApp/models"

type AllCategoryResponse struct {
	ID          int     `json:"id"`
	Description string  `json:"description"`
	Title       string  `json:"title"`
	Price       float32 `json:"price"`
}

type CategoryRequest struct {
	ID          int     `json:"id"`
	Description string  `json:"description"`
	Title       string  `json:"title"`
	Price       float32 `json:"price"`
}

type Category struct {
	Category CategoryRequest `json:"category"`
}

func (categoryRequest Category) ToModelCategory() *models.Category {

	modelCategory := models.Category{
		ID:          categoryRequest.Category.ID,
		Description: categoryRequest.Category.Description,
		Title:       categoryRequest.Category.Title,
		Price:       categoryRequest.Category.Price,
	}

	return &modelCategory
}
