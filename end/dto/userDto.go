package dto

import "App/model"

// TouserMyDto 登录用户Dto
func TouserMyDto(user model.User) model.UserDto {

	return model.UserDto{
		UserID: user.ID,
		Name:   user.Name,
		Super: user.Super,
	}
}