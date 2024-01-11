package models

import "database/sql"

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
