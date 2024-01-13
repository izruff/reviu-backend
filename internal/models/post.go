package models

import "gopkg.in/guregu/null.v3"

type Post struct {
	ID                null.Int    `db:"id" json:"id"`
	Title             null.String `db:"title" json:"title"`
	Content           null.String `db:"content" json:"content"`
	AuthorID          null.Int    `db:"author_id" json:"authorId"`
	TopicID           null.Int    `db:"topic_id" json:"topicId"`
	CreatedAt         null.Time   `db:"created_at" json:"createdAt"`
	UpdatedAt         null.Time   `db:"updated_at" json:"updatedAt"`
	DeletedAt         null.Time   `db:"deleted_at" json:"deletedAt"`
	ReasonForDeletion null.String `db:"reason_for_deletion" json:"reasonForDeletion"`
	ModeratorID       null.Int    `db:"moderator_id" json:"moderatorId"`
}
