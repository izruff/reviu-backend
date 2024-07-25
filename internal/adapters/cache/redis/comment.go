package redis

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/izruff/reviu-backend/internal/core/domain"
)

const (
	commentNamespace = "comment"
)

func (c *RedisCache) GetCommentFieldsByID(id int64, fields ...string) (*domain.Comment, *CacheError) {
	key := joinAsKey(commentNamespace, id)
	fieldMap, err := c.hmget(key, fields...)
	if err != nil {
		return nil, err
	}

	comment := &domain.Comment{}
	for field, val := range fieldMap {
		switch field {
		case "content":
			comment.Content.SetValid(val)
		case "authorId":
			intVal, err := strconv.ParseInt(val, 10, 64)
			if err != nil {
				return nil, newErrInternal(err)
			}
			comment.AuthorID.SetValid(intVal)
		case "postAndParentCommentId":
			parts := strings.SplitN(val, "|", 2)

			postID, err := strconv.ParseInt(parts[0], 10, 64)
			if err != nil {
				return nil, newErrInternal(err)
			}
			comment.PostID.SetValid(postID)

			if parts[1] != "" {
				parentCommentID, err := strconv.ParseInt(parts[1], 10, 64)
				if err != nil {
					return nil, newErrInternal(err)
				}
				comment.ParentCommentID.SetValid(parentCommentID)
			}
		case "topicId":
			intVal, err := strconv.ParseInt(val, 10, 64)
			if err != nil {
				return nil, newErrInternal(err)
			}
			comment.AuthorID.SetValid(intVal)
		case "createdAt":
			timeVal, err := time.Parse(time.RFC3339Nano, val)
			if err != nil {
				return nil, newErrInternal(err)
			}
			comment.CreatedAt.SetValid(timeVal)
		case "updatedAt":
			timeVal, err := time.Parse(time.RFC3339Nano, val)
			if err != nil {
				return nil, newErrInternal(err)
			}
			comment.UpdatedAt.SetValid(timeVal)
		case "deletedDetails":
			if val != "" {
				parts := strings.SplitN(val, "|", 3)

				deletedAt, err := time.Parse(time.RFC3339Nano, parts[0])
				if err != nil {
					return nil, newErrInternal(err)
				}
				comment.UpdatedAt.SetValid(deletedAt)

				moderatorID, err := strconv.ParseInt(parts[1], 10, 64)
				if err != nil {
					return nil, newErrInternal(err)
				}
				comment.AuthorID.SetValid(moderatorID)

				comment.ReasonForDeletion.SetValid(parts[2])
			}
		case "voteCount":
			intVal, err := strconv.ParseInt(val, 10, 64)
			if err != nil {
				return nil, newErrInternal(err)
			}
			comment.VoteCount.SetValid(intVal)
		}
	}

	return comment, nil
}

func (c *RedisCache) SetCommentFieldsByID(id int64, comment *domain.Comment) *CacheError {
	key := joinAsKey(commentNamespace, id)
	fieldMap := make(map[string]string)
	if comment.Content.Valid {
		fieldMap["content"] = comment.Content.String
	}
	if comment.AuthorID.Valid {
		fieldMap["authorId"] = strconv.FormatInt(comment.AuthorID.Int64, 10)
	}
	if comment.PostID.Valid {
		fieldMap["postAndParentCommentId"] = strconv.FormatInt(comment.PostID.Int64, 10) + "|"
		if comment.ParentCommentID.Valid {
			fieldMap["postAndParentCommentId"] += strconv.FormatInt(comment.ParentCommentID.Int64, 10)
		}
	}
	if comment.CreatedAt.Valid {
		fieldMap["createdAt"] = comment.CreatedAt.Time.Format(time.RFC3339Nano)
	}
	if comment.UpdatedAt.Valid {
		fieldMap["updatedAt"] = comment.UpdatedAt.Time.Format(time.RFC3339Nano)
	}
	if comment.DeletedAt.Valid {
		if !comment.ReasonForDeletion.Valid || !comment.ModeratorID.Valid {
			return newErrInternal(errors.New("`reasonForDeletion` and `moderatorId` required alongside `deletedAt`"))
		}
		fieldMap["deletedDetails"] = strings.Join([]string{
			comment.DeletedAt.Time.Format(time.RFC3339Nano),
			strconv.FormatInt(comment.ModeratorID.Int64, 10),
			comment.ReasonForDeletion.String,
		}, "|")
	}
	if comment.VoteCount.Valid {
		fieldMap["voteCount"] = strconv.FormatInt(comment.VoteCount.Int64, 10)
	}

	if err := c.hset(key, fieldMap); err != nil {
		return err
	}
	return nil
}
