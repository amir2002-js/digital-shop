package products

import (
	"github.com/amir2002-js/digital-shop/internal/domain/Comments"
	"github.com/amir2002-js/digital-shop/internal/domain/gallery"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	PriceBeforeOff decimal.Decimal `gorm:"type:numeric(100,2);not null"`
	PriceAfterOff  decimal.Decimal `gorm:"column:price_after_off;->;type:numeric(12,2)"`
	Name           string
	Description    string
	Discount       decimal.Decimal `gorm:"type:numeric(4,2);default:0"`
	CategoryID     *uint
	Stock          uint `gorm:"default:0"`
	//	روابط
	Comments []Comments.Comments `json:"comments" gorm:"foreignKey:ProductID"`
	Tags     []string            `json:"tags" gorm:"many2many:product_tags;"`
	Gallery  []gallery.Gallery   `json:"gallery" gorm:"one2many:product_gallery;"`
}
