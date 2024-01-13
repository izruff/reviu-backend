package models

import "gopkg.in/guregu/null.v3"

type TaggedPost struct {
	PostID null.Int `db:"post_id" json:"postId"`
	TagID  null.Int `db:"tag_id" json:"tagId"`
}
