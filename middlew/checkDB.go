package middlew

import (
	"net/http"

	"viandasApp/db"
)

/* chequeobd es el middwler q me permite conocer el estado de la BD*/
func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		db := db.GetDB()
		if db == nil {
			http.Error(rw, "Conexion perdida con la DB", http.StatusInternalServerError)
			return
		}
		next.ServeHTTP(rw, r)
	}
}
