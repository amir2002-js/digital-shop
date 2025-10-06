package productsUsecase

import (
	"context"
	"github.com/amir2002-js/digital-shop/internal/domain/products"
)

type ProductsUseCase interface {
	Create(ctx context.Context, product *products.Product) (*products.Product, error)
	Update(ctx context.Context, product *products.Product) (*products.Product, error)
	ReadById(ctx context.Context, productID int) (*products.Product, error)
	ReadAll(ctx context.Context) ([]products.Product, error)
	Delete(ctx context.Context, productID int) error
}
