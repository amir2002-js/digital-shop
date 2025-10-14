package repository

import (
	"context"
	"errors"
	"github.com/amir2002-js/digital-shop/internal/domain/gallery"
	"gorm.io/gorm"
)

func (r *GormDb) AddToGallery(ctx context.Context, img *gallery.Gallery) error {
	result := r.DB.WithContext(ctx).Model(&gallery.Gallery{}).Create(img)
	return result.Error
}

func (r *GormDb) RemoveFromGallery(ctx context.Context, id int) error {
	result := r.DB.WithContext(ctx).Model(&gallery.Gallery{}).Where("id = ?", id).Delete(&gallery.Gallery{})
	return result.Error
}

func (r *GormDb) GetImageByProductId(ctx context.Context, id int) ([]gallery.Gallery, error) {
	var imagesArr []gallery.Gallery
	result := r.DB.WithContext(ctx).Model(&gallery.Gallery{}).Where("product_id = ?", id).Find(&imagesArr)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return imagesArr, nil
		}
		return nil, result.Error
	}

	return imagesArr, nil
}
