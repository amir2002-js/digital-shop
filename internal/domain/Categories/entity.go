package Categories

import (
	"github.com/amir2002-js/digital-shop/internal/domain/products"
	"gorm.io/gorm"
)

type Categories struct {
	gorm.Model
	Name     string             `gorm:"unique;not null"`
	Products []products.Product `json:"products" gorm:"foreignKey:CategoryID;constraint:OnDelete:SET NULL;"`
}
