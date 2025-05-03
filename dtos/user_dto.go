package dtos

import "time"

type UserDTO struct {
	ID        uint64     `json:"id"`
	Username  string     `json:"username"`
	Email     string     `json:"email"`
	Avatar    string     `json:"avatar"`
	CreatedAt *time.Time `json:"created_at"`
	Online    bool       `json:"online"`
}

type SimpleUserDTO struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
	Online   bool   `json:"online"`
}
