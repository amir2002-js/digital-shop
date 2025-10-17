package tagsHandler

import (
	cacheService "github.com/amir2002-js/digital-shop/internal/services/cache"
	tagService "github.com/amir2002-js/digital-shop/internal/services/tags"
	"github.com/go-playground/validator/v10"
)

type TagsHandler struct {
	h        *tagService.TagService
	cache    *cacheService.RedisCacheServe
	validate *validator.Validate
}

func NewTagsHandler(h *tagService.TagService, cache *cacheService.RedisCacheServe, v *validator.Validate) *TagsHandler {
	return &TagsHandler{
		h:        h,
		cache:    cache,
		validate: v,
	}
}
