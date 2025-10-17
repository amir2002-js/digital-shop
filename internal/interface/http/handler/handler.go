package handler

import (
	galleryHandler "github.com/amir2002-js/digital-shop/internal/interface/http/handler/gallery"
	productsHandler "github.com/amir2002-js/digital-shop/internal/interface/http/handler/products"
	tagsHandler "github.com/amir2002-js/digital-shop/internal/interface/http/handler/tags"
	"github.com/amir2002-js/digital-shop/internal/interface/http/handler/user"
)

type Handler struct {
	User    *usersHandler.UsersHandler
	Product *productsHandler.ProductsHandler
	Gallery *galleryHandler.GalleryHandler
	Tags    *tagsHandler.TagsHandler
}

func NewHandler(user *usersHandler.UsersHandler, product *productsHandler.ProductsHandler, gallery *galleryHandler.GalleryHandler, tags *tagsHandler.TagsHandler) *Handler {
	return &Handler{User: user, Product: product, Gallery: gallery, Tags: tags}
}
