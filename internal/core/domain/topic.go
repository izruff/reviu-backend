package domain

import "gopkg.in/guregu/null.v3"

type Topic struct {
	ID          null.Int    `db:"id" json:"id"`
	Topic       null.String `db:"topic" json:"topic"`
	Hub         null.String `db:"hub" json:"hub"`
	Description null.String `db:"description" json:"description"`
	CreatedAt   null.Time   `db:"created_at" json:"createdAt"`
}
