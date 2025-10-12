package gallery

import "gorm.io/gorm"

type Gallery struct {
	gorm.Model
	Url       string `gorm:"unique"`
	IsMain    bool   `gorm:"default:false;column:is_main;"`
	ProductId uint   `gorm:"index;column:product_id;"`
}
