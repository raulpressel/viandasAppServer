package middlew

import (
	"net/http"

	"viandasApp/routes"
)

/*ValidoJWT permite validar el JWT que nos viene en la peticion*/

func ValidoJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := routes.ProcessToken(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "Error en el tokeeeeen! "+err.Error(), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	}
}
