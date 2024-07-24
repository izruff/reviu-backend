package ports

import "time"

type Cache interface {
	Set(key string, value interface{}) error
	SetWithExpiry(key string, value interface{}, expiration time.Duration) error
	SetAll(kvMap map[string]interface{}) error
	Get(key string, dest interface{}) error
	DeleteAll(key ...string) error
}
