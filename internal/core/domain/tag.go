package models

import "gopkg.in/guregu/null.v3"

type Tag struct {
	ID        null.Int    `db:"id" json:"id"`
	Tag       null.String `db:"tag" json:"tag"`
	Hub       null.String `db:"hub" json:"hub"`
	CreatedAt null.Time   `db:"created_at" json:"createdAt"`
}
