package models

import "gopkg.in/guregu/null.v3"

type User struct {
	ID           null.Int    `db:"id" json:"id"`
	Email        null.String `db:"email" json:"email"`
	PasswordHash null.String `db:"password_hash" json:"passwordHash"`
	ModRole      null.Bool   `db:"mod_role" json:"modRole"`
	Username     null.String `db:"username" json:"username"`
	Nickname     null.String `db:"nickname" json:"nickname"`
	About        null.String `db:"about" json:"about"`
	CreatedAt    null.Time   `db:"created_at" json:"createdAt"`
}
