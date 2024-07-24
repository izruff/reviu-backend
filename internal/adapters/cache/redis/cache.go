package redis

import (
	"context"
	"encoding/json"
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
	return c.SetWithExpiry(key, value, 0)
}

func (c *RedisCache) SetWithExpiry(key string, value interface{}, expiration time.Duration) error {
	bytes, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return c.client.Set(c.ctx, key, bytes, expiration).Err()
}

func (c *RedisCache) SetAll(kvMap map[string]interface{}) error {
	if len(kvMap) == 0 {
		return nil
	}
	kvMapMarshalled := make(map[string]interface{})
	for k, v := range kvMap {
		bytes, err := json.Marshal(v)
		if err != nil {
			return err
		}
		kvMapMarshalled[k] = bytes
	}
	return c.client.MSet(c.ctx, kvMapMarshalled).Err()
}

func (c *RedisCache) Get(key string, dest interface{}) error {
	str, err := c.client.Get(c.ctx, key).Result()
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(str), dest)
}

func (c *RedisCache) DeleteAll(keys ...string) error {
	return c.client.Del(c.ctx, keys...).Err()
}
