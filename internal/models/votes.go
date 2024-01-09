package models

import "time"

type Vote struct {
	Up        bool      `db:"up" json:"up"`
	PostID    int32     `db:"post_id" json:"postId"`
	UserID    int32     `db:"user_id" json:"userId"`
	CreatedAt time.Time `db:"created_at" json:"createdAt"`
}
