package dtos

import "myproject/forum/models"

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

func ConvertToUserDTO(user *models.User) *UserDTO {
	return &UserDTO{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	}
}

func ConvertToSimpleUserDTO(user *models.User) *SimpleUserDTO {
	return &SimpleUserDTO{
		ID:        user.ID,
		Username:  user.Username,
		AvatarUrl: user.AvatarUrl,
	}
}
