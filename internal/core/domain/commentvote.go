package domain

import "gopkg.in/guregu/null.v3"

type CommentVote struct {
	Up        null.Bool `db:"up" json:"up"`
	CommentID null.Int  `db:"comment_id" json:"commentId"`
	UserID    null.Int  `db:"user_id" json:"userId"`
	CreatedAt null.Time `db:"created_at" json:"createdAt"`
}
