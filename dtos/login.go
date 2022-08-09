package dtos

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

/*RespuestaLogin tiene el tokenm que se devuelve con el login*/
type LoginResponse struct {
	Token string `json:"token"`
}
