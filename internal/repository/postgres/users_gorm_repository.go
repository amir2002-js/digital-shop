package repository

import (
	"context"
	"errors"
	"github.com/amir2002-js/digital-shop/internal/domain/users"
	"gorm.io/gorm"
)

func (r *GormDb) Register(ctx context.Context, user *users.User) (*users.User, error) {
	result := r.DB.WithContext(ctx).Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (r *GormDb) Login(ctx context.Context, email string) (*users.User, error) {
	var ourUser users.User
	result := r.DB.WithContext(ctx).First(&ourUser, "email = ?", email)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, result.Error
	}

	return &ourUser, nil
}

func (r *GormDb) IsEmailExist(ctx context.Context, email string) (*users.User, error) {
	var user users.User
	result := r.DB.WithContext(ctx).First(&user, "email = ?", email)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &user, nil
}
