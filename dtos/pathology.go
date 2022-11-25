package dtos

import "viandasApp/models"

type PathologyRequest struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
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
	}

	return &pathologyModel
}
