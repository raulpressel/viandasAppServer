package dtos

type DiscountRequest struct {
	Discount Discount `json:"discount"`
}

type Discount struct {
	ID          int     `json:"id"`
	Description string  `json:"description"`
	Cant        int     `json:"cant"`
	Percentage  float32 `json:"percentage"`
}

type DiscountResponse struct {
	Discount []Discount `json:"discount"`
}
