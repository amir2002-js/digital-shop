package producttags

type ProductTag struct {
	ProductID int `gorm:"column:product_id;not null"`
	TagID     int `gorm:"column:tag_id;not null"`
}
