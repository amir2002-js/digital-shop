package usecaseUser

import (
	"context"
	"github.com/amir2002-js/digital-shop/internal/domain/users"
)

type UserUseCase interface {
	Register(ctx context.Context, user *users.User) (*users.User, error)
	Login(ctx context.Context, user *users.User) error
}
