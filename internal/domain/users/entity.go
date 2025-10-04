package users

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username   string `json:"username" gorm:"column:username;type:varchar(30);not null"`
	HashedPass string `json:"_" gorm:"column:hashedPass;type:varchar(255);not null"`
	Email      string `json:"email" gorm:"column:email;type:varchar(255);unique;not null"`
	Role       string `json:"-" gorm:"column:role;type:role_type;default:'user'"`
}
