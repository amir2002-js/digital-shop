package cacheRepo

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type RedisCacheRepo struct {
	ClientRepo *redis.Client
}

func NewRedisCacheRepo(client *redis.Client) *RedisCacheRepo {
	return &RedisCacheRepo{ClientRepo: client}
}

func (r *RedisCacheRepo) Set(ctx context.Context, key, value string, ttlSeconds int) error {
	return r.ClientRepo.Set(ctx, key, value, time.Duration(ttlSeconds)*time.Second).Err()
}

func (r *RedisCacheRepo) Get(ctx context.Context, key string) (string, error) {
	return r.ClientRepo.Get(ctx, key).Result()
}

func (r *RedisCacheRepo) Delete(ctx context.Context, key string) error {
	return r.ClientRepo.Del(ctx, key).Err()
}
