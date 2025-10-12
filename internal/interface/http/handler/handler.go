package handler

import (
	galleryHandler "github.com/amir2002-js/digital-shop/internal/interface/http/handler/gallery"
	productsHandler "github.com/amir2002-js/digital-shop/internal/interface/http/handler/products"
	"github.com/amir2002-js/digital-shop/internal/interface/http/handler/user"
)

type Handler struct {
	User    *usersHandler.UsersHandler
	Product *productsHandler.ProductsHandler
	Gallery *galleryHandler.GalleryHandler
}

func NewHandler(user *usersHandler.UsersHandler, product *productsHandler.ProductsHandler, gallery *galleryHandler.GalleryHandler) *Handler {
	return &Handler{User: user, Product: product, Gallery: gallery}
}
