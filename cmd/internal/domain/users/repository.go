package users

type UserRepository interface {
	Register(user *User) (*User , error)
	Login(user *User) error
}