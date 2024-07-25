package ports

import "github.com/izruff/reviu-backend/internal/core/domain"

type Cache interface {
	GetUserFieldsByID(id int64, fields ...string) (*domain.User, *CacheError)
	SetUserFieldsByID(id int64, user *domain.User) *CacheError
	GetUserIDByUsername(username string) (int64, *CacheError)
	SetUserIDByUsername(username string, id int64) *CacheError

	GetTopicFieldsByID(id int64, fields ...string) (*domain.Topic, *CacheError)
	SetTopicFieldsByID(id int64, topic *domain.Topic) *CacheError
	GetTopicIDByTopic(topic string) (int64, *CacheError)
	SetTopicIDByTopic(topic string, id int64) *CacheError

	GetPostFieldsByID(id int64, fields ...string) (*domain.Post, *CacheError)
	SetPostFieldsByID(id int64, post *domain.Post) *CacheError

	GetCommentFieldsByID(id int64, fields ...string) (*domain.Comment, *CacheError)
	SetCommentFieldsByID(id int64, comment *domain.Comment) *CacheError

	GetTagFieldsByID(id int64, fields ...string) (*domain.Tag, *CacheError)
	SetTagFieldsByID(id int64, tag *domain.Tag) *CacheError
}

// TODO: CacheError
type CacheError struct {
	Err error
}
