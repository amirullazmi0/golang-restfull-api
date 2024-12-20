package service

import (
	"amirullazmi/belajar-restfull-api/helper"
	"amirullazmi/belajar-restfull-api/model/domain"
	"amirullazmi/belajar-restfull-api/model/web"
	"amirullazmi/belajar-restfull-api/repository"
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func (service *UserServiceImpl) Create(ctx context.Context, request web.UserCreateRequest) web.UserResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user := domain.User{
		Name:  request.Name,
		Email: request.Email,
	}

	user = service.UserRepository.Save(ctx, tx, user)

	return helper.ToUserResponse(user)
}
func (service *UserServiceImpl) Update(ctx context.Context, request web.UserUpdateRequest) web.UserResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindById(ctx, tx, request.ID)

	helper.PanicIfError(err)

	user = domain.User{
		ID:    request.ID,
		Name:  request.Name,
		Email: request.Email,
	}

	user = service.UserRepository.Save(ctx, tx, user)

	return helper.ToUserResponse(user)

}
func (service *UserServiceImpl) Delete(ctx context.Context, userId string) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindById(ctx, tx, userId)

	helper.PanicIfError(err)

	service.UserRepository.Delete(ctx, tx, user)

}
func (service *UserServiceImpl) FindById(ctx context.Context, userId string) web.UserResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindById(ctx, tx, userId)

	helper.PanicIfError(err)

	return helper.ToUserResponse(user)

}
func (service *UserServiceImpl) FindAll(ctx context.Context) []web.UserResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	users := service.UserRepository.FindAll(ctx, tx)

	var userResponses []web.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, helper.ToUserResponse(user))
	}

	return userResponses
}
