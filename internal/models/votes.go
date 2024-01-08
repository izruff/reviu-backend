package models

import "time"

type Vote struct {
	Up        bool      `db:"up"`
	PostID    int32     `db:"post_id"`
	UserID    int32     `db:"user_id"`
	CreatedAt time.Time `db:"created_at"`
}
