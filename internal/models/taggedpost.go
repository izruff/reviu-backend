package models

type TaggedPost struct {
	PostID int32 `db:"post_id"`
	TagID  int32 `db:"tag_id"`
}
