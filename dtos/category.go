package dtos

import "viandasApp/models"

type AllCategoryResponse struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Title       string `json:"title"`
	Price       string `json:"price"`
}

type CategoryRequest struct {
	ID          int     `json:"id"`
	Description string  `json:"description"`
	Title       string  `json:"title"`
	Price       float32 `json:"price"`
}

type CategoryDeleteRequest struct {
	ID int `json:"id"`
}

func (categoryRequest CategoryRequest) ToModelCategory() *models.Category {

	modelCategory := models.Category{
		ID:          categoryRequest.ID,
		Description: categoryRequest.Description,
		Title:       categoryRequest.Title,
		Price:       categoryRequest.Price,
	}

	return &modelCategory
}
