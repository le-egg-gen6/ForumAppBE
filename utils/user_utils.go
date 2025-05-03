package utils

import (
	"forum/dtos"
	"forum/models"
)

func ConvertToUserDTO(user *models.User) *dtos.UserDTO {
	if user == nil {
		return nil
	}
	userDTO := &dtos.UserDTO{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}
	if user.Avatar != nil {
		userDTO.Avatar = user.Avatar.URL
	}
	userDTO.Online = IsUserOnline(user.ID)
	return userDTO
}

func ConvertToSimpleUserDTO(user *models.User) *dtos.SimpleUserDTO {
	if user == nil {
		return nil
	}
	simpleUserDTO := &dtos.SimpleUserDTO{
		ID:       user.ID,
		Username: user.Username,
	}
	if user.Avatar != nil {
		simpleUserDTO.Avatar = user.Avatar.URL
	}
	simpleUserDTO.Online = IsUserOnline(user.ID)
	return simpleUserDTO
}
