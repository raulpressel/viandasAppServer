package middlew

import (
	"net/http"
	"viandasApp/handlers"
)

func ValidateJWTAdmin(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

		_, admin, err := handlers.ProcessToken(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(rw, "Token incorrecto "+err.Error(), http.StatusBadRequest)
			return
		}
		if admin {
			next.ServeHTTP(rw, r)

		} else {
			http.Error(rw, "Forbidden", http.StatusForbidden)
			return
		}

	})
}
