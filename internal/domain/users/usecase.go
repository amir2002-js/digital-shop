package users

import "context"

type UserRepository interface {
	Register(ctx context.Context,user *User) (*User , error)
	Login(ctx context.Context,user *User) error
}