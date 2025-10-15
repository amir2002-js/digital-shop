package usecase

import (
	"context"
	"github.com/amir2002-js/digital-shop/internal/domain/products"
	"github.com/amir2002-js/digital-shop/internal/domain/productsTags"
)

type ProductsUseCase interface {
	Create(ctx context.Context, product *products.Product) error
	Update(ctx context.Context, product *products.Product) error
	ReadById(ctx context.Context, productID int) (*products.Product, error)
	ReadAll(ctx context.Context) ([]products.Product, error)
	Delete(ctx context.Context, productID int) error
	AddToTags(ctx context.Context, tags []producttags.ProductTag) error
	FindTag(ctx context.Context, tagID int) (bool, error)
}
