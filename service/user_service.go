package service

import (
	"amirullazmi/belajar-restfull-api/model/web"
	"context"
)

type CategoryService interface {
	Create(ctx context.Context, request web.UserCreateRequest) web.UserResponse
	Update(ctx context.Context, request web.UserUpdateRequest) web.UserResponse
	Delete(ctx context.Context, UserId string)
	FindById(ctx context.Context, UserId string) web.UserResponse
	FindAll(ctx context.Context) []web.UserResponse
}
