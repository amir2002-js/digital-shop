package tagService

import (
	"context"
	"github.com/amir2002-js/digital-shop/internal/domain/Tags"
	"github.com/amir2002-js/digital-shop/internal/usecase"
)

type TagService struct {
	repo usecase.TagsUseCase
}

func NewTagService(repo usecase.TagsUseCase) *TagService {
	return &TagService{repo: repo}
}

func (serve *TagService) CreateTag(ctx context.Context, name string) (*Tags.Tag, error) {
	return serve.repo.CreateTag(ctx, name)
}

func (serve *TagService) ReadByIdTag(ctx context.Context, id int) (*Tags.Tag, error) {
	return serve.repo.ReadByIdTag(ctx, id)
}

func (serve *TagService) ReadAllTag(ctx context.Context) ([]Tags.Tag, error) {
	return serve.repo.ReadAllTag(ctx)
}

func (serve *TagService) DeleteByIdTag(ctx context.Context, id int) error {
	return serve.repo.DeleteByIdTag(ctx, id)
}
