package dtos

type UserDTO struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SimpleUserDTO struct {
	ID        uint64 `json:"id"`
	Username  string `json:"username"`
	AvatarUrl string `json:"avatarUrl"`
}
