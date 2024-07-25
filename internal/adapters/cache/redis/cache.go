package redis

import (
	"context"
	"errors"
	"strconv"
	"strings"
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

func (c *RedisCache) hset(key string, fieldMap map[string]string) *CacheError {
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

func (c *RedisCache) hmget(key string, fields ...string) (map[string]string, *CacheError) {
	// Retrieves all fields; return error if any is missing
	fieldValues, err := c.client.HMGet(c.ctx, key, fields...).Result()
	if err != nil {
		return nil, newErrRedis(err)
	}
	result := make(map[string]string, len(fields))
	for i, field := range fields {
		if fieldValues[i] == nil {
			return nil, newErrRedis(errors.New("some fields not found"))
		}
		result[field] = fieldValues[i].(string) // go-redis guarantees this is a string
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

func joinAsKey(names ...interface{}) string {
	strNames := make([]string, len(names))
	for i, name := range names {
		switch v := name.(type) {
		case int64:
			strNames[i] = strconv.FormatInt(v, 10)
		case string:
			strNames[i] = v
		default:
			// Since joinAsKey is called internally with known types, this should not happen
			panic("unsupported type for Redis key naming")
		}
	}
	return strings.Join(strNames, ":")
}
