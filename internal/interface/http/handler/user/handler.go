package usersHandler

import (
	usersServices "github.com/amir2002-js/digital-shop/internal/services/users"
	"github.com/go-playground/validator/v10"
)

type UsersHandler struct {
	h        *usersServices.UsersServices
	validate *validator.Validate
}

func NewUsersHandler(userServe *usersServices.UsersServices) *UsersHandler {
	return &UsersHandler{
		h:        userServe,
		validate: validator.New(),
	}
}
