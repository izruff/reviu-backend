package repository

import "github.com/izruff/reviu-backend/internal/models"

func (q *PostgresQueries) CreateVote(newVote *models.Vote) error {
	// TODO: error handling when form is incomplete or already exists
	if _, err := q.create("votes", []string{"up", "post_id", "user_id"}, false, newVote); err != nil {
		return err
	}

	return nil
}

func (q *PostgresQueries) GetVotesFromPostID(postID int64) ([]*models.Vote, error) {
	return nil, nil // TODO
}

func (q *PostgresQueries) CountVotesFromPostID(postID int64) (int64, error) {
	return 0, nil // TODO
}

func (q *PostgresQueries) UpdateVote(postID int64, userID int64) error {
	return nil // TODO
}

func (q *PostgresQueries) DeleteVote(postID int64, userID int64) error {
	return nil // TODO
}
