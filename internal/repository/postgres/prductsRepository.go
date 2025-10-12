package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/amir2002-js/digital-shop/internal/domain/products"
	"gorm.io/gorm"
)

func (r *GormDb) ReadAll(ctx context.Context) ([]products.Product, error) {
	var allProducts []products.Product
	result := r.DB.WithContext(ctx).Model(&products.Product{}).Find(&allProducts)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return allProducts, errors.New("products not available")
		}
		return nil, result.Error
	}
	return allProducts, nil
}

func (r *GormDb) ReadById(ctx context.Context, id int) (*products.Product, error) {
	var product products.Product
	result := r.DB.WithContext(ctx).Model(&product).Where("id = ?", id).Preload("Comments").Preload("Gallery").Find(&product)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			errContent := fmt.Sprintf("Product with id %d not found", id)
			return nil, errors.New(errContent)
		}
		return nil, result.Error
	}

	return &product, nil
}

func (r *GormDb) Create(ctx context.Context, product *products.Product) error {
	result := r.DB.WithContext(ctx).Model(&products.Product{}).Create(product)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *GormDb) Update(ctx context.Context, product *products.Product) error {
	result := r.DB.WithContext(ctx).Model(&products.Product{}).Save(product)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *GormDb) Delete(ctx context.Context, id int) error {
	result := r.DB.WithContext(ctx).Model(&products.Product{}).Delete(&products.Product{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
