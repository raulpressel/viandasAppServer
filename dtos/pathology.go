package dtos

import "viandasApp/models"

type PathologyRequest struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Color       string `json:"color"`
}

type PathologyResponse struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Checked     bool   `json:"checked"`
}

type Pathology struct {
	Pathology PathologyRequest `json:"pathology"`
}

func (pathologyRequest Pathology) ToModelPathology() *models.Pathology {

	pathologyModel := models.Pathology{
		ID:          pathologyRequest.Pathology.ID,
		Description: pathologyRequest.Pathology.Description,
		Color:       pathologyRequest.Pathology.Color,
	}

	return &pathologyModel
}
