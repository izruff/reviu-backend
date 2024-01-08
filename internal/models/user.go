package models

import (
	"database/sql"
	"time"
)

type User struct {
	ID           int32          `db:"id"`
	Email        string         `db:"email"`
	PasswordHash string         `db:"passwordHash"`
	ModRole      bool           `db:"modRole"`
	Username     string         `db:"username"`
	Nickname     sql.NullString `db:"nickname"`
	About        sql.NullString `db:"about"`
	CreatedAt    time.Time      `db:"createdAt"`
}
