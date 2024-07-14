package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisCache struct {
	client *redis.Client
	ctx    context.Context
}

func NewRedisCache(client *redis.Client) *RedisCache {
	return &RedisCache{
		client: client,
		ctx:    context.Background(),
	}
}

func (c *RedisCache) Set(key string, value interface{}) error {
	return c.client.Set(c.ctx, key, value, 0).Err()
}

func (c *RedisCache) SetWithExpiry(key string, value interface{}, expiration time.Duration) error {
	return c.client.Set(c.ctx, key, value, expiration).Err()
}

func (c *RedisCache) SetAll(kvmap map[string]interface{}) error {
	return c.client.MSet(c.ctx, kvmap).Err()
}

func (c *RedisCache) Get(key string) (string, error) {
	return c.client.Get(c.ctx, key).Result()
}

func (c *RedisCache) GetAll(keys ...string) ([]interface{}, error) {
	return c.client.MGet(c.ctx, keys...).Result()
}

func (c *RedisCache) DeleteAll(keys ...string) error {
	return c.client.Del(c.ctx, keys...).Err()
}
