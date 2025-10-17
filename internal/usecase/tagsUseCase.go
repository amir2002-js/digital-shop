package usecase

import (
	"context"
	"github.com/amir2002-js/digital-shop/internal/domain/Tags"
)

type TagsUseCase interface {
	CreateTag(ctx context.Context, name string) (*Tags.Tag, error)
	ReadByIdTag(ctx context.Context, id int) (*Tags.Tag, error)
	ReadAllTag(ctx context.Context) ([]Tags.Tag, error)
	DeleteByIdTag(ctx context.Context, id int) error
}
