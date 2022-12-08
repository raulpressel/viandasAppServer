package dtos

import "viandasApp/models"

type Address struct {
	Address               AddressRequest `json:"address"`
	IDClient              int            `json:"idClient"`
	IDNewFavouriteAddress int            `json:"idNewFavouriteAddress"`
	IDOldFavouriteAddress int            `json:"idOldFavouriteAddress"`
}

type AddressRequest struct {
	ID          int    `json:"id"`
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
	Favourite   bool            `json:"favourite"`
	City        AllCityResponse `json:"city"`
}

func (addressRequest Address) ToModelAddress() *models.Address {

	addressModel := models.Address{
		ID:          addressRequest.Address.ID,
		Street:      addressRequest.Address.Street,
		Floor:       addressRequest.Address.Floor,
		Number:      addressRequest.Address.Number,
		Departament: addressRequest.Address.Departament,
		Observation: addressRequest.Address.Observation,
	}

	return &addressModel
}
