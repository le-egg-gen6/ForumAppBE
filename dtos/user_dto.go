package dtos

import "time"

type UserDTO struct {
	ID           uint      `json:"id"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	Avatar       string    `json:"avatar"`
	CreatedAt    time.Time `json:"createdAt"`
	Online       bool      `json:"online"`
	FriendStatus int       `json:"friendStatus"`
}

type SimpleUserDTO struct {
	ID           uint   `json:"id"`
	Username     string `json:"username"`
	Avatar       string `json:"avatar"`
	Online       bool   `json:"online"`
	FriendStatus int    `json:"friendStatus"`
}
