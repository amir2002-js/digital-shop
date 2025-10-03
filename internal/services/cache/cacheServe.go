package cacheService

import (
	"context"
	cacheRepo "github.com/amir2002-js/digital-shop/internal/repository/cache"
)

type RedisCacheServe struct {
	ClientServ *cacheRepo.RedisCacheRepo
}

func NewRedisCacheService(client *cacheRepo.RedisCacheRepo) *RedisCacheServe {
	return &RedisCacheServe{ClientServ: client}
}

func (r *RedisCacheServe) Set(ctx context.Context, key, value string, ttlSeconds int) error {
	return r.ClientServ.Set(ctx, key, value, ttlSeconds)
}

func (r *RedisCacheServe) Get(ctx context.Context, key string) (string, error) {
	return r.ClientServ.Get(ctx, key)
}

func (r *RedisCacheServe) Delete(ctx context.Context, key string) error {
	return r.ClientServ.Delete(ctx, key)
}
