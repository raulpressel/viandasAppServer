package middlew

import (
	"net/http"

	"viandasApp/handlers"
)

/*ValidoJWT permite validar el JWT que nos viene en la peticion*/

func ValidateJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, err := handlers.ProcessToken(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "Error en el tokeeeeen! "+err.Error(), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	}
}
