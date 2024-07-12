package postgres

import (
	"github.com/izruff/reviu-backend/internal/models"
	"gopkg.in/guregu/null.v3"
)

func (q *PostgresQueries) CreatePostVote(newVote *models.PostVote) error {
	// TODO: error handling when form is incomplete or already exists
	if _, err := q.create("post_votes", []string{"up", "post_id", "user_id"}, false, newVote); err != nil {
		return err
	}

	return nil
}

func (q *PostgresQueries) GetPostVoteValue(postID int64, userID int64) (*null.Bool, error) {
	var vote models.PostVote
	if err := q.selectOne(&vote, "post_votes", "up", "post_id=$1 AND user_id=$2", postID, userID); err != nil {
		voted := null.NewBool(false, false)
		return &voted, nil // TODO: error handling for other internal errors
	}

	return &vote.Up, nil
}

func (q *PostgresQueries) GetVotesFromPostID(postID int64) ([]models.PostVote, error) {
	votes := []models.PostVote{}
	if err := q.selectAll(&votes, "post_votes", "*", "WHERE post_id=:post_id", "", postID); err != nil {
		return nil, err
	}

	return votes, nil
}

func (q *PostgresQueries) CountVotesFromPostID(postID int64) (int64, int64, error) {
	upCount, err := q.count("post_votes", "up", "up=t AND post_id=$1", postID)
	if err != nil {
		return 0, 0, err
	}

	downCount, err := q.count("post_votes", "up", "up=f AND post_id=$1", postID)
	if err != nil {
		return 0, 0, err
	}

	return upCount, downCount, nil
}

func (q *PostgresQueries) UpdatePostVote(up bool, postID int64, userID int64) error {
	updatedVote := &models.PostVote{
		Up:     null.NewBool(up, true),
		PostID: null.NewInt(postID, true),
		UserID: null.NewInt(userID, true),
	}
	if err := q.updateByPK("post_votes", []string{"up"}, []string{"post_id", "user_id"}, updatedVote); err != nil {
		return err // TODO: error handling when vote does not exist
	}

	return nil
}

func (q *PostgresQueries) DeletePostVote(postID int64, userID int64) error {
	if err := q.deleteWhere("post_votes", true, "post_id=$1 AND user_id=$2", postID, userID); err != nil {
		return err // TODO: error handling when vote does not exist
	}

	return nil
}
