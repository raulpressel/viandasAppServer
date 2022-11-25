package dtos

type Address struct {
	Direction AddressRequest `json:"address"`
}

type AddressRequest struct {
	Street      string `json:"street"`
	Number      string `json:"number"`
	Floor       string `json:"floor"`
	Departament string `json:"departament"`
	Observation string `json:"observation"`
}

type AddressRespone struct {
	ID          int             `json:"id"`
	Street      string          `json:"street"`
	Number      string          `json:"number"`
	Floor       string          `json:"floor"`
	Departament string          `json:"departament"`
	Observation string          `json:"observation"`
	City        AllCityResponse `json:"city"`
}
