package productsHandler

import (
	cacheService "github.com/amir2002-js/digital-shop/internal/services/cache"
	"github.com/go-playground/validator/v10"
)

type ProductsHandler struct {
	cache    *cacheService.RedisCacheServe
	validate *validator.Validate
}

func NewProductsHandler(cache *cacheService.RedisCacheServe, v *validator.Validate) *ProductsHandler {
	return &ProductsHandler{
		cache:    cache,
		validate: v,
	}
}
