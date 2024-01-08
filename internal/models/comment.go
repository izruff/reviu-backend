package models

import (
	"database/sql"
	"time"
)

type Comment struct {
	ID                int32          `db:"id"`
	Content           sql.NullString `db:"content"`
	AuthorID          int32          `db:"authorId"`
	PostID            int32          `db:"postId"`
	ParentCommentID   sql.NullInt32  `db:"parentCommentId"`
	CreatedAt         time.Time      `db:"createdAt"`
	UpdatedAt         sql.NullTime   `db:"updatedAt"`
	DeletedAt         sql.NullTime   `db:"deletedAt"`
	ReasonForDeletion sql.NullString `db:"reasonForDeletion"`
	ModeratorID       sql.NullInt32  `db:"moderatorId"`
}
