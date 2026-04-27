package dtos

type AllProductResponse struct {
	ID                uint    `json:"id"`
	Title             string  `json:"title"`
	Description       string  `json:"description"`
	Price             float32 `json:"price"`
	Available         bool    `json:"available"`
	ProductCategoryID int     `json:"productCategoryId"`
}

type ProductRequest struct {
	Product struct {
		ID                uint    `json:"id"`
		Title             string  `json:"title"`
		Description       string  `json:"description"`
		Price             float32 `json:"price"`
		Available         bool    `json:"available"`
		ProductCategoryID int     `json:"productCategoryId"`
	} `json:"product"`
}
