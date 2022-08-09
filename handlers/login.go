package handlers

import (
	"encoding/json"
	"net/http"
	"time"
	"viandasApp/db"
	"viandasApp/dtos"
	"viandasApp/jwt"
)

/* login realiza el login */

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	//var t models.User

	var loginReq dtos.LoginRequest

	err := json.NewDecoder(r.Body).Decode(&loginReq)
	if err != nil {
		http.Error(w, "Decode error"+err.Error(), 400)
		return
	}
	/* 	if len(loginReq.Email) == 0 {
		http.Error(w, "email requerido si o si", 400)
		return
	} */

	user, exist := db.GetLogin(loginReq.Email, loginReq.Password)

	if !exist {
		http.Error(w, "User or Pass incorrectos", 400)
		return
	}

	jwtKey, err := jwt.GenerateJWT(user)
	if err != nil {
		http.Error(w, "ocurrio un error al intentar generar el token correspondiente"+err.Error(), 400)
		return
	}

	resp := dtos.LoginResponse{
		Token: jwtKey,
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})
}
