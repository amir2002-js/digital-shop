package repository

import (
	"context"
	"errors"
	"github.com/amir2002-js/digital-shop/internal/domain/Tags"
	"gorm.io/gorm"
)

func (r *GormDb) CreateTag(ctx context.Context, name string) (*Tags.Tag, error) {
	var tag Tags.Tag
	tag.Name = name
	result := r.DB.WithContext(ctx).Model(&Tags.Tag{}).Create(&tag)
	if result.Error != nil {
		return nil, result.Error
	}
	return &tag, nil
}

func (r *GormDb) ReadByIdTag(ctx context.Context, id int) (*Tags.Tag, error) {
	var tag Tags.Tag
	result := r.DB.WithContext(ctx).Model(&Tags.Tag{}).Where("id = ?", id).First(&tag)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	return &tag, nil
}

func (r *GormDb) ReadAllTag(ctx context.Context) ([]Tags.Tag, error) {
	var tags []Tags.Tag
	result := r.DB.WithContext(ctx).Model(&Tags.Tag{}).Find(&tags)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return tags, nil
}

func (r *GormDb) DeleteByIdTag(ctx context.Context, id int) error {
	result := r.DB.WithContext(ctx).Model(&Tags.Tag{}).Where("id = ?", id).Delete(&Tags.Tag{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
