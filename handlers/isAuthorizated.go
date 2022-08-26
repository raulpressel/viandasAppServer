package handlers

import (
	"encoding/json"
	"net/http"
)

type AuthorizationResponse struct {
	Validate bool `json:"resp"`
}

func IsAuthorizated(w http.ResponseWriter, r *http.Request){
		
	
	authorizationResponseModel := AuthorizationResponse{}
	authorizationResponseModel.Validate = true
	_, _, err := ProcessToken(r.Header.Get("Authorization"))
		
	if err != nil {
		authorizationResponseModel.Validate = false
		} 
	
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(authorizationResponseModel)
}