package productsService

import (
	"context"
	"github.com/amir2002-js/digital-shop/internal/domain/products"
	productsUsecase "github.com/amir2002-js/digital-shop/internal/usecase"
)

type ProductsService struct {
	Repo productsUsecase.ProductsUseCase
}

func NewProductsService(repo productsUsecase.ProductsUseCase) *ProductsService {
	return &ProductsService{
		Repo: repo,
	}
}

func (serve *ProductsService) ReadById(ctx context.Context, productId int) (*products.Product, error) {
	return serve.Repo.ReadById(ctx, productId)
}

func (serve *ProductsService) Create(ctx context.Context, product *products.Product) error {
	return serve.Repo.Create(ctx, product)
}

func (serve *ProductsService) Update(ctx context.Context, product *products.Product) error {
	return serve.Repo.Update(ctx, product)
}

func (serve *ProductsService) Delete(ctx context.Context, productId int) error {
	return serve.Repo.Delete(ctx, productId)
}

func (serve *ProductsService) ReadAll(ctx context.Context) ([]products.Product, error) {
	return serve.Repo.ReadAll(ctx)
}
