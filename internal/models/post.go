package models

import (
	"database/sql"
	"time"
)

type Post struct {
	ID                int32          `db:"id"`
	Title             sql.NullString `db:"title"`
	Content           sql.NullString `db:"content"`
	AuthorID          int32          `db:"authorId"`
	TopicID           int32          `db:"topicId"`
	CreatedAt         time.Time      `db:"createdAt"`
	UpdatedAt         sql.NullTime   `db:"updatedAt"`
	DeletedAt         sql.NullTime   `db:"deletedAt"`
	ReasonForDeletion sql.NullString `db:"reasonForDeletion"`
	ModeratorID       sql.NullInt32  `db:"moderatorId"`
}
