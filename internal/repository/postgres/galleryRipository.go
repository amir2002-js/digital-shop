package repository

import (
	"context"
	"github.com/amir2002-js/digital-shop/internal/domain/gallery"
)

func (r *GormDb) AddToGallery(ctx context.Context, img *gallery.Gallery) error {
	result := r.DB.WithContext(ctx).Model(&gallery.Gallery{}).Create(img)
	return result.Error
}

func (r *GormDb) RemoveFromGallery(ctx context.Context, id int) error {
	result := r.DB.WithContext(ctx).Model(&gallery.Gallery{}).Where("id = ?", id).Delete(&gallery.Gallery{})
	return result.Error
}
