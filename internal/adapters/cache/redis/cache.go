package redis

import (
	"context"
	"time"

	"github.com/izruff/reviu-backend/internal/core/ports"
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

type CacheError = ports.CacheError

// Internal server error; for any unexpected error that is not categorized here
func newErrInternal(err error) *CacheError {
	return &CacheError{
		Err: err,
	}
}

// Error returned by Redis
// TODO: Need to break this down further
func newErrRedis(err error) *CacheError {
	return &CacheError{
		Err: err,
	}
}

func (c *RedisCache) set(key string, value interface{}) *CacheError {
	return c.setWithExpiry(key, value, 0)
}

func (c *RedisCache) setWithExpiry(key string, value interface{}, expiration time.Duration) *CacheError {
	err := c.client.Set(c.ctx, key, value, expiration).Err()
	if err != nil {
		return newErrRedis(err)
	}
	return nil
}

func (c *RedisCache) hset(key string, fieldMap map[string]interface{}) *CacheError {
	err := c.client.HSet(c.ctx, key, fieldMap).Err()
	if err != nil {
		return newErrRedis(err)
	}
	return nil
}

func (c *RedisCache) get(key string) (string, *CacheError) {
	value, err := c.client.Get(c.ctx, key).Result()
	if err != nil {
		return "", newErrRedis(err)
	}
	return value, nil
}

func (c *RedisCache) hmget(key string, fields ...string) (map[string]interface{}, *CacheError) {
	fieldValues, err := c.client.HMGet(c.ctx, key, fields...).Result()
	if err != nil {
		return nil, newErrRedis(err)
	}
	result := make(map[string]interface{}, len(fields))
	for i, field := range fields {
		result[field] = fieldValues[i]
	}
	return result, nil
}

func (c *RedisCache) deleteAll(keys ...string) *CacheError {
	err := c.client.Del(c.ctx, keys...).Err()
	if err != nil {
		return newErrRedis(err)
	}
	return nil
}
