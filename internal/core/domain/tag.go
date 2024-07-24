package domain

import "gopkg.in/guregu/null.v3"

type Tag struct {
	ID        null.Int    `db:"id" json:"id"`
	Tag       null.String `db:"tag" json:"tag"`
	CreatedAt null.Time   `db:"created_at" json:"createdAt"`
}
