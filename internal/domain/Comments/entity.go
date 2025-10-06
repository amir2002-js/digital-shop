package Comments

import (
	"gorm.io/gorm"
)

type Comments struct {
	gorm.Model
	ProductID   uint   `gorm:"index;not null"`
	UserID      uint   `gorm:"index;not null"`
	Title       string `gorm:"not null"`
	Description string `gorm:"not null"`
}
