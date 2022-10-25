package middlew

import (
	"net/http"
	"viandasApp/handlers"
)

func ValidateJWTAdmin(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {

		_, admin, err := handlers.ProcessToken(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(rw, "Error en el tokeeeeen! "+err.Error(), http.StatusBadRequest)
			return
		}
		if !admin {
			//http.Error(rw, "Unauthorized", http.StatusInternalServerError)
			return
		}
		next.ServeHTTP(rw, r)

	}
}
