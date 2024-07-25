package ports

import "github.com/izruff/reviu-backend/internal/core/domain"

// TODO: Create caches for paginated results
type Cache interface {
	GetUserByID(id int64) (*domain.User, *CacheError)
	SetUserByID(id int64, updatedUser *domain.User) *CacheError
	GetUserIDByUsername(username string) (int64, *CacheError)
	SetUserIDByUsername(username string, id int64) *CacheError

	GetTopicByID(id int64) (*domain.Topic, *CacheError)
	SetTopicByID(id int64, updatedTopic *domain.Topic) *CacheError

	GetPostByID(id int64) (*domain.Post, *CacheError)
	SetPostByID(id int64, updatedPost *domain.Post) *CacheError
	DeletePostByID(id int64) *CacheError

	GetCommentByID(id int64) (*domain.Comment, *CacheError)
	SetCommentByID(id int64, updatedComment *domain.Comment) *CacheError
	DeleteCommentByID(id int64) *CacheError

	GetTagByID(id int64) (*domain.Tag, *CacheError)
}

// TODO: CacheError
type CacheError struct {
	Err error
}
