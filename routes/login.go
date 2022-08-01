package routes

import (
	"encoding/json"
	"net/http"
	"time"
	"viandasApp/db"
	"viandasApp/jwt"
	"viandasApp/models"
)

/* login realiza el login */

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	var t models.User

	
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "User or Pass incorrectos"+err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "email requerido si o si", 400)
		return
	}

	documento, existe := db.TryLogin(t.Email, t.Password)

	if !existe {
		http.Error(w, "User or Pass incorrectos", 400)
		return
	}

	jwtKey, err := jwt.GeneroJWT(documento)
	if err != nil {
		http.Error(w, "ocurrio un error al intentar generar el token correspondiente"+err.Error(), 400)
		return
	}

	resp := models.RespuestaLogin{
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
