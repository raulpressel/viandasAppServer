package dtos

import "time"

type Client struct {
	Client RegisterRequest `json:"client"`
}

type RegisterRequest struct {
	ID             int                `json:"id"`
	PhonePrimary   string             `json:"phonePrimary"`
	PhoneSecondary string             `json:"phoneSecondary"`
	ObsClient      string             `json:"observation"`
	BornDate       string             `json:"bornDate"`
	Pathologies    []PathologyRequest `json:"pathologies"`
	Address        []AddressRequest   `json:"addresses"`
}

type ClientRespone struct {
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
