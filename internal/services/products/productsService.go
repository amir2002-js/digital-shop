package productsService

import (
	"context"
	"github.com/amir2002-js/digital-shop/internal/domain/products"
	producttags "github.com/amir2002-js/digital-shop/internal/domain/productsTags"
	"github.com/amir2002-js/digital-shop/internal/usecase"
)

type ProductsService struct {
	Repo usecase.ProductsUseCase
}

func NewProductsService(repo usecase.ProductsUseCase) *ProductsService {
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

func (serve *ProductsService) AddToTags(ctx context.Context, tags []producttags.ProductTag) error {
	return serve.Repo.AddToTags(ctx, tags)
}

func (serve *ProductsService) FindTag(ctx context.Context, tagID int) (bool, error) {
	return serve.Repo.FindTag(ctx, tagID)
}
