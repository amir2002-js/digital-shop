package usersHandler

import (
	cacheService "github.com/amir2002-js/digital-shop/internal/services/cache"
	usersServices "github.com/amir2002-js/digital-shop/internal/services/users"
	"github.com/go-playground/validator/v10"
)

type UsersHandler struct {
	h        *usersServices.UsersServices
	cache    *cacheService.RedisCacheServe
	validate *validator.Validate
}

func NewUsersHandler(userServe *usersServices.UsersServices, cache *cacheService.RedisCacheServe, validate *validator.Validate) *UsersHandler {
	return &UsersHandler{
		h:        userServe,
		cache:    cache,
		validate: validate,
	}
}
