package users

import (
	"github.com/amir2002-js/digital-shop/internal/domain/Comments"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username   string              `json:"username" gorm:"column:username;type:varchar(30);not null"`
	HashedPass string              `json:"-" gorm:"column:hashed_pass;type:varchar(255);not null"`
	Email      string              `json:"email" gorm:"column:email;type:varchar(255);unique;not null"`
	Role       string              `json:"-" gorm:"column:role;type:varchar(20);default:'user'"`
	Comments   []Comments.Comments `json:"comments" gorm:"foreignKey:UserID"`
}
