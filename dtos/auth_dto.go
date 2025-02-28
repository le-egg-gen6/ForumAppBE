package dtos

type AuthDTO struct {
	Token string `json:"token"`
}

type LoginDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterDTO struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
