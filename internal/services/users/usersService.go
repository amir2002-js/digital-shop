package usersServices

import (
	"context"
	"github.com/amir2002-js/digital-shop/internal/domain/users"
	usecaseUser "github.com/amir2002-js/digital-shop/internal/usecase"
)

type UsersServices struct {
	Repo usecaseUser.UserUseCase
}

func NewUsersServices(repo usecaseUser.UserUseCase) *UsersServices {
	return &UsersServices{
		Repo: repo,
	}
}

func (serve *UsersServices) Register(ctx context.Context, user *users.User) (*users.User, error) {
	return serve.Repo.Register(ctx, user)
}

func (serve *UsersServices) Login(ctx context.Context, user *users.User) error {
	return serve.Repo.Login(ctx, user)
}
