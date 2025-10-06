package products

import (
	"github.com/amir2002-js/digital-shop/internal/domain/Comments"
	"github.com/amir2002-js/digital-shop/internal/domain/Tags"
	"github.com/lib/pq"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	PriceBeforeOff decimal.Decimal `gorm:"type:numeric(100,2);not null"`
	PriceAfterOff  decimal.Decimal `gorm:"type:numeric(100,2);not null"`
	Name           string
	Description    string
	ImgURL         pq.StringArray  `gorm:"type:text[];not null;default:'{}'"`
	Off            decimal.Decimal `gorm:"type:numeric(4,2);default:0"`
	CategoryID     *uint
	Stock          uint `gorm:"default:0"`
	//	روابط
	Comments []Comments.Comments `json:"comments" gorm:"foreignKey:ProductID"`
	Tags     []Tags.Tag          `json:"tags" gorm:"many2many:product_tags;"`
}
