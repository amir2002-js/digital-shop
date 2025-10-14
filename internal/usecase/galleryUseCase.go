package usecase

import (
	"context"
	"github.com/amir2002-js/digital-shop/internal/domain/gallery"
)

type GalleryUseCase interface {
	AddToGallery(ctx context.Context, image *gallery.Gallery) error
	RemoveFromGallery(ctx context.Context, id int) error
	GetImageByProductId(ctx context.Context, id int) ([]gallery.Gallery, error)
}
