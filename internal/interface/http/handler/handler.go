package handler

import (
	usersHandler "github.com/amir2002-js/digital-shop/internal/interface/http/handler/user"
)

type Handler struct {
	User *usersHandler.UsersHandler
}

func NewHandler(user *usersHandler.UsersHandler) *Handler {
	return &Handler{User: user}
}
