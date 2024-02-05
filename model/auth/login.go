package auth

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token  string `json:"token"`
	Expire string `json:"expire"`
}
