package handler

import (
	usersHandler "github.com/amir2002-js/digital-shop/internal/interface/http/handler/user"
)

type Handler struct {
	userHandler *usersHandler.UsersHandler
}

func NewHandler(userHandler *usersHandler.UsersHandler) *Handler {
	return &Handler{userHandler: userHandler}
}
