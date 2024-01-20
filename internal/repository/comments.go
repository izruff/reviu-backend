package repository

import (
	"errors"
	"time"

	"github.com/izruff/reviu-backend/internal/models"
	"gopkg.in/guregu/null.v3"
)

func (q *PostgresQueries) CreateComment(newComment *models.Comment) (int64, error) {
	// TODO: error handling when form is incomplete
	postID, err := q.create("comments", []string{"content", "author_id", "post_id", "parent_comment_id"}, true, newComment)
	if err != nil {
		return 0, err
	}

	return postID, nil
}

func (q *PostgresQueries) GetCommentByID(id int64) (*models.Comment, error) {
	comment := &models.Comment{}
	if err := q.selectOne(comment, "comments", "*", "id=$1", id); err != nil {
		return nil, err // TODO: error handling when comment does not exist
	}

	return comment, nil
}

// TODO: implement filters and preferences
func (q *PostgresQueries) GetCommentsWithOptions(options interface{}) ([]*models.Comment, error) {
	return nil, nil
}

func (q *PostgresQueries) UpdateCommentByID(updatedComment *models.Comment) error {
	if !updatedComment.ID.Valid {
		return errors.New("ID not provided")
	}

	var columns []string
	if updatedComment.Content.Valid {
		columns = append(columns, "content")
	}
	// TODO: error handling if nothing is updated

	updatedComment.UpdatedAt = null.NewTime(time.Now(), true)
	columns = append(columns, "updated_at")

	if err := q.updateByID("comments", columns, updatedComment); err != nil {
		return err // TODO: error handling when comment does not exist
	}

	return nil
}

func (q *PostgresQueries) MarkCommentAsDeletedByID(id int64, reason string, moderatorID int64) error {
	updatedComment := &models.Comment{
		ID:                null.NewInt(id, true),
		ReasonForDeletion: null.NewString(reason, true),
		ModeratorID:       null.NewInt(moderatorID, true),
		DeletedAt:         null.NewTime(time.Now(), true),
	}
	if err := q.updateByID("comments", []string{"reason_for_deletion", "moderator_id", "deleted_at"}, updatedComment); err != nil {
		return err // TODO: error handling when comment does not exist
	}

	return nil
}
