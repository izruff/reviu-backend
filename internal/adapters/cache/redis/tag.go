package redis

import (
	"time"

	"github.com/izruff/reviu-backend/internal/core/domain"
)

const (
	tagNamespace = "tag"
)

func (c *RedisCache) GetTagFieldsByID(id int64, fields ...string) (*domain.Tag, *CacheError) {
	key := joinAsKey(tagNamespace, id)
	fieldMap, err := c.hmget(key, fields...)
	if err != nil {
		return nil, err
	}

	tag := &domain.Tag{}
	for field, val := range fieldMap {
		switch field {
		case "tag":
			tag.Tag.SetValid(val)
		case "createdAt":
			timeVal, err := time.Parse(time.RFC3339Nano, val)
			if err != nil {
				return nil, newErrInternal(err)
			}
			tag.CreatedAt.SetValid(timeVal)
		}
	}

	return tag, nil
}

func (c *RedisCache) SetTagFieldsByID(id int64, tag *domain.Tag) *CacheError {
	key := joinAsKey(tagNamespace, id)
	fieldMap := make(map[string]string)
	if tag.Tag.Valid {
		fieldMap["tag"] = tag.Tag.String
	}
	if tag.CreatedAt.Valid {
		fieldMap["createdAt"] = tag.CreatedAt.Time.Format(time.RFC3339Nano)
	}

	if err := c.hset(key, fieldMap); err != nil {
		return err
	}
	return nil
}
