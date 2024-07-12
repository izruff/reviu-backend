package models

import "gopkg.in/guregu/null.v3"

type PostView struct {
	PostID    null.Int  `db:"post_id" json:"postId"`
	UserID    null.Int  `db:"user_id" json:"userId"`
	CreatedAt null.Time `db:"created_at" json:"createdAt"`
}
