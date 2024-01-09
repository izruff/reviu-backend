package models

import "time"

type Topic struct {
	ID        int32     `db:"id" json:"id"`
	Topic     string    `db:"topic" json:"topic"`
	Hub       string    `db:"hub" json:"hub"`
	CreatedAt time.Time `db:"created_at" json:"createdAt"`
}
