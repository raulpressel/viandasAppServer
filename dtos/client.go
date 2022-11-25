package dtos

import "time"

type RegisterRequest struct {
	Client Register `json:"client"`
}

type ClientResponse struct {
	Client Client `json:"client"`
}

type Register struct {
	ID             int                `json:"id"`
	PhonePrimary   string             `json:"phonePrimary"`
	PhoneSecondary string             `json:"phoneSecondary"`
	ObsClient      string             `json:"observation"`
	BornDate       string             `json:"bornDate"`
	Pathologies    []PathologyRequest `json:"pathologies"`
	Address        []AddressRequest   `json:"addresses"`
}

type Client struct {
	ID             int                 `json:"id"`
	Name           string              `json:"name"`
	LastName       string              `json:"lastName"`
	Email          string              `json:"email"`
	PhonePrimary   string              `json:"phonePrimary"`
	PhoneSecondary string              `json:"phoneSecondary"`
	ObsClient      string              `json:"observation"`
	BornDate       time.Time           `json:"bornDate"`
	Pathologies    []PathologyResponse `json:"pathologies"`
	Address        []AddressRespone    `json:"addresses"`
}
