package ports

import "time"

type Cache interface {
	Set(key string, value interface{}) error
	SetWithExpiry(key string, value interface{}, expiration time.Duration) error
	SetAll(kvmap map[string]interface{}) error
	Get(key string) (string, error)
	DeleteAll(key ...string) error
}
