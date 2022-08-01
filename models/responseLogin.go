package models

/*RespuestaLogin tiene el tokenm que se devuelve con el login*/
type RespuestaLogin struct {
	Token string `json:"token"`
}
