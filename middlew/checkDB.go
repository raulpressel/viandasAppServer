package middlew

import (
	"net/http"

	"viandasApp/db"
)

/* chequeobd es el middwler q me permite conocer el estado de la BD*/
func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if db.CheckConnection() == 0 {
			http.Error(w, "Conexion perdida con la DB", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
