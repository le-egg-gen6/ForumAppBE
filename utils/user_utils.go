package utils

import (
	"forum/dtos"
	"forum/models"
)

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
	return simpleUserDTO
}
