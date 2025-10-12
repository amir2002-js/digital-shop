package galleryHandler

import (
	cacheService "github.com/amir2002-js/digital-shop/internal/services/cache"
	galleryService "github.com/amir2002-js/digital-shop/internal/services/gallery"
	"github.com/go-playground/validator/v10"
)

type GalleryHandler struct {
	h        *galleryService.GalleryService
	cache    *cacheService.RedisCacheServe
	validate *validator.Validate
}

func NewGalleryHandler(h *galleryService.GalleryService, cache *cacheService.RedisCacheServe, v *validator.Validate) *GalleryHandler {
	return &GalleryHandler{
		h:        h,
		cache:    cache,
		validate: v,
	}
}
