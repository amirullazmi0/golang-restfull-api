package helper

import (
	"amirullazmi/belajar-restfull-api/model/domain"
	"amirullazmi/belajar-restfull-api/model/web"
)

func ToUserResponse(user domain.User) web.UserResponse {
	return web.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}
