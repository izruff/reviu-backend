package models

import "time"

type Topic struct {
	ID        int32     `db:"id"`
	Topic     string    `db:"topic"`
	Hub       string    `db:"hub"`
	CreatedAt time.Time `db:"createdAt"`
}
