package postgres

import (
	"github.com/izruff/reviu-backend/internal/core/domain"
	"gopkg.in/guregu/null.v3"
)

func (r *PostgresRepository) CreateCommentVote(newVote *domain.CommentVote) error {
	// TODO: error handling when form is incomplete or already exists
	if _, err := r.create("comment_votes", []string{"up", "comment_id", "user_id"}, false, newVote); err != nil {
		return err
	}

	return nil
}

func (r *PostgresRepository) GetCommentVoteValue(commentID int64, userID int64) (*null.Bool, error) {
	var vote domain.CommentVote
	if err := r.selectOne(&vote, "comment_votes", "up", "comment_id=$1 AND user_id=$2", commentID, userID); err != nil {
		voted := null.NewBool(false, false)
		return &voted, nil // TODO: error handling for other internal errors
	}

	return &vote.Up, nil
}

func (r *PostgresRepository) GetVotesFromCommentID(commentID int64) ([]domain.CommentVote, error) {
	votes := []domain.CommentVote{}
	if err := r.selectAll(&votes, "comment_votes", "*", "WHERE comment_id=:comment_id", "", commentID); err != nil {
		return nil, err
	}

	return votes, nil
}

func (r *PostgresRepository) CountVotesFromCommentID(commentID int64) (int64, int64, error) {
	upCount, err := r.count("comment_votes", "up", "up=t AND comment_id=$1", commentID)
	if err != nil {
		return 0, 0, err
	}

	downCount, err := r.count("comment_votes", "up", "up=f AND comment_id=$1", commentID)
	if err != nil {
		return 0, 0, err
	}

	return upCount, downCount, nil
}

func (r *PostgresRepository) UpdateCommentVote(up bool, commentID int64, userID int64) error {
	updatedVote := &domain.CommentVote{
		Up:        null.NewBool(up, true),
		CommentID: null.NewInt(commentID, true),
		UserID:    null.NewInt(userID, true),
	}
	if err := r.updateByPK("comment_votes", []string{"up"}, []string{"comment_id", "user_id"}, updatedVote); err != nil {
		return err // TODO: error handling when vote does not exist
	}

	return nil
}

func (r *PostgresRepository) DeleteCommentVote(commentID int64, userID int64) error {
	if err := r.deleteWhere("comment_votes", true, "comment_id=$1 AND user_id=$2", commentID, userID); err != nil {
		return err // TODO: error handling when vote does not exist
	}

	return nil
}
