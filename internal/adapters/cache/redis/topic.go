package redis

import (
	"strconv"
	"time"

	"github.com/izruff/reviu-backend/internal/core/domain"
)

const (
	topicNamespace          = "topic"
	topicIDByTopicNamespace = "topic:name"
)

func (c *RedisCache) GetTopicFieldsByID(id int64, fields ...string) (*domain.Topic, *CacheError) {
	key := joinAsKey(topicNamespace, id)
	fieldMap, err := c.hmget(key, fields...)
	if err != nil {
		return nil, err
	}

	topic := &domain.Topic{}
	for field, val := range fieldMap {
		switch field {
		case "topic":
			topic.Topic.SetValid(val)
		case "description":
			topic.Description.SetValid(val)
		case "createdAt":
			timeVal, err := time.Parse(time.RFC3339Nano, val)
			if err != nil {
				return nil, newErrInternal(err)
			}
			topic.CreatedAt.SetValid(timeVal)
		}
	}

	return topic, nil
}

func (c *RedisCache) SetTopicFieldsByID(id int64, topic *domain.Topic) *CacheError {
	key := joinAsKey(topicNamespace, id)
	fieldMap := make(map[string]string)
	if topic.Topic.Valid {
		fieldMap["topic"] = topic.Topic.String
	}
	if topic.Description.Valid {
		fieldMap["description"] = topic.Description.String
	}
	if topic.CreatedAt.Valid {
		fieldMap["createdAt"] = topic.CreatedAt.Time.Format(time.RFC3339Nano)
	}

	if err := c.hset(key, fieldMap); err != nil {
		return err
	}
	return nil
}

func (c *RedisCache) GetTopicIDByTopic(topic string) (int64, *CacheError) {
	key := joinAsKey(topicIDByTopicNamespace, topic)
	strID, cErr := c.get(key)
	if cErr != nil {
		return 0, cErr
	}

	ID, err := strconv.ParseInt(strID, 10, 64)
	if err != nil {
		return 0, newErrInternal(err)
	}
	return ID, nil
}

func (c *RedisCache) SetTopicIDByTopic(topic string, id int64) *CacheError {
	key := joinAsKey(topicIDByTopicNamespace, topic)
	if err := c.set(key, id); err != nil {
		return err
	}
	return nil
}
