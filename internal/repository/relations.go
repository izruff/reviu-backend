package repository

import "github.com/izruff/reviu-backend/internal/models"

func (q *PostgresQueries) CreateRelation(newRelation *models.Relation) error {
	// TODO: error handling when form is incomplete or already exists
	if _, err := q.create("relations", []string{"follower_id", "following_id"}, false, newRelation); err != nil {
		return err
	}

	return nil
}

func (q *PostgresQueries) GetFollowersFromUserID(userID int64) ([]*models.Relation, error) {
	return nil, nil // TODO
}

func (q *PostgresQueries) CountFollowersFromPostID(userID int64) (int64, error) {
	return 0, nil // TODO
}

func (q *PostgresQueries) GetFollowingsFromUserID(userID int64) ([]*models.Relation, error) {
	return nil, nil // TODO
}

func (q *PostgresQueries) CountFollowingsFromPostID(userID int64) (int64, error) {
	return 0, nil // TODO
}

func (q *PostgresQueries) DeleteRelation(followerID int64, followingID int64) error {
	return nil // TODO
}
