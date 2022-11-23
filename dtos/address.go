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
