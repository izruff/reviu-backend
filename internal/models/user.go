package models

import (
	"database/sql"
	"time"
)

type User struct {
	ID           int32          `db:"id" json:"id"`
	Email        string         `db:"email" json:"email"`
	PasswordHash string         `db:"password_hash" json:"passwordHash"`
	ModRole      bool           `db:"mod_role" json:"modRole"`
	Username     string         `db:"username" json:"username"`
	Nickname     sql.NullString `db:"nickname" json:"nickname"`
	About        sql.NullString `db:"about" json:"about"`
	CreatedAt    time.Time      `db:"created_at" json:"createdAt"`
}
