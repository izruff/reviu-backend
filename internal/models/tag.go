package models

import "time"

type Tag struct {
	ID        int32     `db:"id" json:"id"`
	Tag       string    `db:"tag" json:"tag"`
	Hub       string    `db:"hub" json:"hub"`
	CreatedAt time.Time `db:"created_at" json:"createdAt"`
}
