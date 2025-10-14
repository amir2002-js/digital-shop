package galleryService

import (
	"context"
	"github.com/amir2002-js/digital-shop/internal/domain/gallery"
	"github.com/amir2002-js/digital-shop/internal/usecase"
)

type GalleryService struct {
	Repo usecase.GalleryUseCase
}

func NewGalleryService(repo usecase.GalleryUseCase) *GalleryService {
	return &GalleryService{
		Repo: repo,
	}
}

func (g *GalleryService) AddToGallery(ctx context.Context, img *gallery.Gallery) error {
	return g.Repo.AddToGallery(ctx, img)
}

func (g *GalleryService) RemoveFromGallery(ctx context.Context, id int) error {
	return g.Repo.RemoveFromGallery(ctx, id)
}

func (g *GalleryService) GetImageByProductId(ctx context.Context, id int) ([]gallery.Gallery, error) {
	return g.Repo.GetImageByProductId(ctx, id)
}
