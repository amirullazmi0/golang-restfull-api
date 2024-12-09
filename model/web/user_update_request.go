package web

type UserUpdateRequest struct {
	ID    string `validate:"required"`
	Name  string
	Email string
}
