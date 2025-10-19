package usecase

import (
	"context"
	"github.com/amir2002-js/digital-shop/internal/domain/users"
)

type UserUseCase interface {
	Register(ctx context.Context, user *users.User) (*users.User, error)
	Login(ctx context.Context, email string) (*users.User, error)
	IsEmailExist(ctx context.Context, email string) (*users.User, error)
	AddToBasket(ctx context.Context, userID, ProductID int) error
	RemoveFromBasket(ctx context.Context, userID, ProductID int) error
}
