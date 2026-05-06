package dtos

type AllProductCategoryResponse struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type ProductCategoryRequest struct {
	ProductCategory struct {
		ID          uint   `json:"id"`
		Title       string `json:"title"`
		Description string `json:"description"`
	} `json:"productCategory"`
}
