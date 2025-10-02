package repository

import "gorm.io/gorm"

type GormDb struct {
	DB *gorm.DB
}

func NewGormDb(db *gorm.DB) *GormDb {
	return &GormDb{db}
}
