package users

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username   string `gorm:"column:username;type:varchar(30);not null"`
	HashedPass string `gorm:"column:hashedPass;type:varchar(255);not null"`
	Email      string `gorm:"column:email;type:varchar(255);unique;not null"`
	Role       string `gorm:"column:role;type:role_type;default:'user'"`
}