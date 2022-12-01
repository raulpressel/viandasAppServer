package dtos

type OrderRequest struct {
	IDClient         int                `json:"idClient"`
	Observation      string             `json:"observation"`
	Total            float32            `json:"total"`
	Date             string             `json:"date"`
	DaysOrderRequest []DaysOrderRequest `json:"daysOrderRequest"`
}

type DaysOrderRequest struct {
	Amount      int    `json:"cant"`
	IDDayFood   int    `json:"idDayFood"`
	IDAddress   int    `json:"idAddress"`
	Observation string `json:"observation"`
}
