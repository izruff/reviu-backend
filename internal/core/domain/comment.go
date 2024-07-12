package domain

import "gopkg.in/guregu/null.v3"

type Comment struct {
	ID                null.Int    `db:"id" json:"id"`
	Content           null.String `db:"content" json:"content"`
	AuthorID          null.Int    `db:"author_id" json:"authorId"`
	PostID            null.Int    `db:"post_id" json:"postId"`
	ParentCommentID   null.Int    `db:"parent_comment_id" json:"parentCommentId"`
	CreatedAt         null.Time   `db:"created_at" json:"createdAt"`
	UpdatedAt         null.Time   `db:"updated_at" json:"updatedAt"`
	DeletedAt         null.Time   `db:"deleted_at" json:"deletedAt"`
	ReasonForDeletion null.String `db:"reason_for_deletion" json:"reasonForDeletion"`
	ModeratorID       null.Int    `db:"moderator_id" json:"moderatorId"`
	VoteCount         null.Int    `db:"vote_count" json:"voteCount"`
}
