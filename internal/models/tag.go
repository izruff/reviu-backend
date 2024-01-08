package models

import "time"

type Tag struct {
	ID        int32     `db:"id"`
	Tag       string    `db:"tag"`
	Hub       string    `db:"hub"`
	CreatedAt time.Time `db:"createdAt"`
}
