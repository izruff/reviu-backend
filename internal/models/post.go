package models

import (
	"database/sql"
	"time"
)

type Post struct {
	ID                int32          `db:"id" json:"id"`
	Title             sql.NullString `db:"title" json:"title"`
	Content           sql.NullString `db:"content" json:"content"`
	AuthorID          int32          `db:"author_id" json:"authorId"`
	TopicID           int32          `db:"topic_id" json:"topicId"`
	CreatedAt         time.Time      `db:"created_at" json:"createdAt"`
	UpdatedAt         sql.NullTime   `db:"updated_at" json:"updatedAt"`
	DeletedAt         sql.NullTime   `db:"deleted_at" json:"deletedAt"`
	ReasonForDeletion sql.NullString `db:"reason_for_deletion" json:"reasonForDeletion"`
	ModeratorID       sql.NullInt32  `db:"moderator_id" json:"moderatorId"`
}
