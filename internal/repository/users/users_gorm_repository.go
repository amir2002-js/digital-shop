package uersGORMRepository

import (
	"context"
	"github.com/amir2002-js/digital-shop/internal/domain/users"
)

type UsersServices struct {
	Repo users.UserRepository
}

func NewUsersServices(repo *users.UserRepository) *UsersServices {
	return &UsersServices{
		Repo: repo,
	}
}

func (serve *UsersServices) Register(ctx context.Context, user *users.User) (*users.User, error)  {
	return serve.Repo.Register(ctx ,user)
}