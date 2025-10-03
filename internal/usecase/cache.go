package cacheUsecase

import "context"

type CacheRepository interface {
	Set(ctx context.Context, key string, value string, ttlSeconds int) error
	Get(ctx context.Context, key string) (string, error)
	Delete(ctx context.Context, key string) error
}
