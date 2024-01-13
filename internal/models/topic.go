package models

import "gopkg.in/guregu/null.v3"

type Topic struct {
	ID        null.Int    `db:"id" json:"id"`
	Topic     null.String `db:"topic" json:"topic"`
	Hub       null.String `db:"hub" json:"hub"`
	CreatedAt null.Time   `db:"created_at" json:"createdAt"`
}
