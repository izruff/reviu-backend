package redis

import (
	"strconv"
	"time"

	"github.com/izruff/reviu-backend/internal/core/domain"
)

const (
	userNamespace             = "user"
	userIDByUsernameNamespace = "user:username"
)

func (c *RedisCache) GetUserFieldsByID(id int64, fields ...string) (*domain.User, *CacheError) {
	key := joinAsKey(userNamespace, id)
	fieldMap, err := c.hmget(key, fields...)
	if err != nil {
		return nil, err
	}

	user := &domain.User{}
	for field, val := range fieldMap {
		switch field {
		case "email":
			user.Email.SetValid(val)
		case "modRole":
			user.ModRole.SetValid(val == "true")
		case "username":
			user.Username.SetValid(val)
		case "nickname":
			user.Nickname.SetValid(val)
		case "about":
			user.About.SetValid(val)
		case "createdAt":
			timeVal, err := time.Parse(time.RFC3339Nano, val)
			if err != nil {
				return nil, newErrInternal(err)
			}
			user.CreatedAt.SetValid(timeVal)
		case "rating":
			intVal, err := strconv.ParseInt(val, 10, 64)
			if err != nil {
				return nil, newErrInternal(err)
			}
			user.Rating.SetValid(intVal)
		case "followerCount":
			intVal, err := strconv.ParseInt(val, 10, 64)
			if err != nil {
				return nil, newErrInternal(err)
			}
			user.FollowerCount.SetValid(intVal)
		case "followingCount":
			intVal, err := strconv.ParseInt(val, 10, 64)
			if err != nil {
				return nil, newErrInternal(err)
			}
			user.FollowingCount.SetValid(intVal)
		case "postCount":
			intVal, err := strconv.ParseInt(val, 10, 64)
			if err != nil {
				return nil, newErrInternal(err)
			}
			user.PostCount.SetValid(intVal)
		}
	}

	return user, nil
}

func (c *RedisCache) SetUserFieldsByID(id int64, user *domain.User) *CacheError {
	key := joinAsKey(userNamespace, id)
	fieldMap := make(map[string]string)
	if user.ModRole.Valid {
		if user.ModRole.Bool {
			fieldMap["modRole"] = "true"
		} else {
			fieldMap["modRole"] = "false"
		}
	}
	if user.Username.Valid {
		fieldMap["username"] = user.Username.String
	}
	if user.Nickname.Valid {
		fieldMap["nickname"] = user.Nickname.String
	}
	if user.About.Valid {
		fieldMap["about"] = user.About.String
	}
	if user.CreatedAt.Valid {
		fieldMap["createdAt"] = user.CreatedAt.Time.Format(time.RFC3339Nano)
	}
	if user.Rating.Valid {
		fieldMap["rating"] = strconv.FormatInt(user.Rating.Int64, 10)
	}
	if user.FollowerCount.Valid {
		fieldMap["followerCount"] = strconv.FormatInt(user.FollowerCount.Int64, 10)
	}
	if user.FollowingCount.Valid {
		fieldMap["followingCount"] = strconv.FormatInt(user.FollowingCount.Int64, 10)
	}
	if user.PostCount.Valid {
		fieldMap["postCount"] = strconv.FormatInt(user.PostCount.Int64, 10)
	}

	if err := c.hset(key, fieldMap); err != nil {
		return err
	}
	return nil
}

func (c *RedisCache) GetUserIDByUsername(username string) (int64, *CacheError) {
	key := joinAsKey(userIDByUsernameNamespace, username)
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

func (c *RedisCache) SetUserIDByUsername(username string, id int64) *CacheError {
	key := joinAsKey(userIDByUsernameNamespace, username)
	if err := c.set(key, id); err != nil {
		return err
	}
	return nil
}
