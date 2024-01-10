package models

import (
	"database/sql"
	"errors"
	"time"

	"github.com/izruff/reviu-backend/internal/utils"
)

type User struct {
	ID           sql.NullInt32  `db:"id" json:"id"`
	Email        sql.NullString `db:"email" json:"email"`
	PasswordHash sql.NullString `db:"password_hash" json:"passwordHash"`
	ModRole      sql.NullBool   `db:"mod_role" json:"modRole"`
	Username     sql.NullString `db:"username" json:"username"`
	Nickname     sql.NullString `db:"nickname" json:"nickname"`
	About        sql.NullString `db:"about" json:"about"`
	CreatedAt    sql.NullTime   `db:"created_at" json:"createdAt"`
}

func NewUser(userMap map[string]interface{}) (*User, error) {
	user := &User{
		ModRole: sql.NullBool{Bool: false, Valid: true},
	}

	for key, value := range userMap {
		// TODO: input validation and refactoring
		switch key {
		case "id":
			v, ok := value.(int32)
			if !ok {
				return nil, NewErrInvalidValue(key, "invalid type")
			}
			user.ID.Int32 = v
			user.ID.Valid = true
		case "email":
			v, ok := value.(string)
			if !ok {
				return nil, NewErrInvalidValue(key, "invalid type")
			}
			user.Email.String = v
			user.Email.Valid = true
		case "password":
			v, ok := value.(string)
			if !ok {
				return nil, NewErrInvalidValue(key, "invalid type")
			}
			hash, err := utils.GetPasswordHash(v)
			if err != nil {
				return nil, err
			}
			user.PasswordHash.String = hash
			user.PasswordHash.Valid = true
		case "passwordHash":
			v, ok := value.(string)
			if !ok {
				return nil, NewErrInvalidValue(key, "invalid type")
			}
			user.PasswordHash.String = v
			user.PasswordHash.Valid = true
		case "modRole":
			v, ok := value.(bool)
			if !ok {
				return nil, NewErrInvalidValue(key, "invalid type")
			}
			user.ModRole.Bool = v
			user.ModRole.Valid = true
		case "username":
			v, ok := value.(string)
			if !ok {
				return nil, NewErrInvalidValue(key, "invalid type")
			}
			user.Username.String = v
			user.Username.Valid = true
		case "nickname":
			v, ok := value.(string)
			if !ok {
				return nil, NewErrInvalidValue(key, "invalid type")
			}
			user.Nickname.String = v
			user.Nickname.Valid = true
		case "about":
			v, ok := value.(string)
			if !ok {
				return nil, NewErrInvalidValue(key, "invalid type")
			}
			user.About.String = v
			user.About.Valid = true
		case "createdAt":
			v, ok := value.(time.Time)
			if !ok {
				return nil, NewErrInvalidValue(key, "invalid type")
			}
			user.CreatedAt.Time = v
			user.CreatedAt.Valid = true
		default:
			return nil, errors.New("property does not exist: " + key)
		}
	}

	return user, nil
}
