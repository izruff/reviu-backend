package postgres

import (
	"github.com/izruff/reviu-backend/internal/models"
	"gopkg.in/guregu/null.v3"
)

func (q *PostgresQueries) CreateCommentVote(newVote *models.CommentVote) error {
	// TODO: error handling when form is incomplete or already exists
	if _, err := q.create("comment_votes", []string{"up", "comment_id", "user_id"}, false, newVote); err != nil {
		return err
	}

	return nil
}

func (q *PostgresQueries) GetCommentVoteValue(commentID int64, userID int64) (*null.Bool, error) {
	var vote models.CommentVote
	if err := q.selectOne(&vote, "comment_votes", "up", "comment_id=$1 AND user_id=$2", commentID, userID); err != nil {
		voted := null.NewBool(false, false)
		return &voted, nil // TODO: error handling for other internal errors
	}

	return &vote.Up, nil
}

func (q *PostgresQueries) GetVotesFromCommentID(commentID int64) ([]models.CommentVote, error) {
	votes := []models.CommentVote{}
	if err := q.selectAll(&votes, "comment_votes", "*", "WHERE comment_id=:comment_id", "", commentID); err != nil {
		return nil, err
	}

	return votes, nil
}

func (q *PostgresQueries) CountVotesFromCommentID(commentID int64) (int64, int64, error) {
	upCount, err := q.count("comment_votes", "up", "up=t AND comment_id=$1", commentID)
	if err != nil {
		return 0, 0, err
	}

	downCount, err := q.count("comment_votes", "up", "up=f AND comment_id=$1", commentID)
	if err != nil {
		return 0, 0, err
	}

	return upCount, downCount, nil
}

func (q *PostgresQueries) UpdateCommentVote(up bool, commentID int64, userID int64) error {
	updatedVote := &models.CommentVote{
		Up:        null.NewBool(up, true),
		CommentID: null.NewInt(commentID, true),
		UserID:    null.NewInt(userID, true),
	}
	if err := q.updateByPK("comment_votes", []string{"up"}, []string{"comment_id", "user_id"}, updatedVote); err != nil {
		return err // TODO: error handling when vote does not exist
	}

	return nil
}

func (q *PostgresQueries) DeleteCommentVote(commentID int64, userID int64) error {
	if err := q.deleteWhere("comment_votes", true, "comment_id=$1 AND user_id=$2", commentID, userID); err != nil {
		return err // TODO: error handling when vote does not exist
	}

	return nil
}
