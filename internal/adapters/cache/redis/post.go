package redis

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/izruff/reviu-backend/internal/core/domain"
)

const (
	postNamespace = "post"
)

func (c *RedisCache) GetPostFieldsByID(id int64, fields ...string) (*domain.Post, *CacheError) {
	key := joinAsKey(postNamespace, id)
	fieldMap, err := c.hmget(key, fields...)
	if err != nil {
		return nil, err
	}

	post := &domain.Post{}
	for field, val := range fieldMap {
		switch field {
		case "title":
			post.Title.SetValid(val)
		case "content":
			post.Content.SetValid(val)
		case "authorId":
			intVal, err := strconv.ParseInt(val, 10, 64)
			if err != nil {
				return nil, newErrInternal(err)
			}
			post.AuthorID.SetValid(intVal)
		case "topicId":
			intVal, err := strconv.ParseInt(val, 10, 64)
			if err != nil {
				return nil, newErrInternal(err)
			}
			post.AuthorID.SetValid(intVal)
		case "createdAt":
			timeVal, err := time.Parse(time.RFC3339Nano, val)
			if err != nil {
				return nil, newErrInternal(err)
			}
			post.CreatedAt.SetValid(timeVal)
		case "updatedAt":
			timeVal, err := time.Parse(time.RFC3339Nano, val)
			if err != nil {
				return nil, newErrInternal(err)
			}
			post.UpdatedAt.SetValid(timeVal)
		case "deletedDetails":
			if val != "" {
				parts := strings.SplitN(val, "|", 3)

				deletedAt, err := time.Parse(time.RFC3339Nano, parts[0])
				if err != nil {
					return nil, newErrInternal(err)
				}
				post.UpdatedAt.SetValid(deletedAt)

				moderatorID, err := strconv.ParseInt(parts[1], 10, 64)
				if err != nil {
					return nil, newErrInternal(err)
				}
				post.AuthorID.SetValid(moderatorID)

				post.ReasonForDeletion.SetValid(parts[2])
			}
		case "viewCount":
			intVal, err := strconv.ParseInt(val, 10, 64)
			if err != nil {
				return nil, newErrInternal(err)
			}
			post.ViewCount.SetValid(intVal)
		case "voteCount":
			intVal, err := strconv.ParseInt(val, 10, 64)
			if err != nil {
				return nil, newErrInternal(err)
			}
			post.VoteCount.SetValid(intVal)
		}
	}

	return post, nil
}

func (c *RedisCache) SetPostFieldsByID(id int64, post *domain.Post) *CacheError {
	key := joinAsKey(postNamespace, id)
	fieldMap := make(map[string]string)
	if post.Title.Valid {
		fieldMap["title"] = post.Title.String
	}
	if post.Content.Valid {
		fieldMap["content"] = post.Content.String
	}
	if post.AuthorID.Valid {
		fieldMap["authorId"] = strconv.FormatInt(post.AuthorID.Int64, 10)
	}
	if post.TopicID.Valid {
		fieldMap["topicId"] = strconv.FormatInt(post.TopicID.Int64, 10)
	}
	if post.CreatedAt.Valid {
		fieldMap["createdAt"] = post.CreatedAt.Time.Format(time.RFC3339Nano)
	}
	if post.UpdatedAt.Valid {
		fieldMap["updatedAt"] = post.UpdatedAt.Time.Format(time.RFC3339Nano)
	}
	if post.DeletedAt.Valid {
		if !post.ReasonForDeletion.Valid || !post.ModeratorID.Valid {
			return newErrInternal(errors.New("`reasonForDeletion` and `moderatorId` required alongside `deletedAt`"))
		}
		fieldMap["deletedDetails"] = strings.Join([]string{
			post.DeletedAt.Time.Format(time.RFC3339Nano),
			strconv.FormatInt(post.ModeratorID.Int64, 10),
			post.ReasonForDeletion.String,
		}, "|")
	}
	if post.ViewCount.Valid {
		fieldMap["viewCount"] = strconv.FormatInt(post.ViewCount.Int64, 10)
	}
	if post.VoteCount.Valid {
		fieldMap["voteCount"] = strconv.FormatInt(post.VoteCount.Int64, 10)
	}

	if err := c.hset(key, fieldMap); err != nil {
		return err
	}
	return nil
}
