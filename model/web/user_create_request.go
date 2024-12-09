package web

type UserCreateRequest struct {
	Name  string `validate:"required"`
	Email string `validate:"required"`
}
