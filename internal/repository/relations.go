package repository

import "github.com/izruff/reviu-backend/internal/models"

func (q *PostgresQueries) CreateRelation(newRelation *models.Relation) error {
	// TODO: error handling when form is incomplete or already exists
	if _, err := q.create("relations", []string{"follower_id", "following_id"}, false, newRelation); err != nil {
		return err
	}

	return nil
}

func (q *PostgresQueries) GetFollowersFromUserID(userID int64) ([]models.Relation, error) {
	var relations []models.Relation
	if err := q.selectAll(&relations, "relations", "follower_id", "following_id=$1", "", userID); err != nil {
		return nil, err
	}

	return relations, nil
}

func (q *PostgresQueries) CountFollowersFromUserID(userID int64) (int64, error) {
	// TODO: possibly optimize this by creating a new column in the database
	count, err := q.count("relations", "follower_id", "following_id=$1", userID)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (q *PostgresQueries) GetFollowingsFromUserID(userID int64) ([]models.Relation, error) {
	var relations []models.Relation
	if err := q.selectAll(&relations, "relations", "following_id", "follower_id=$1", "", userID); err != nil {
		return nil, err
	}

	return relations, nil
}

func (q *PostgresQueries) CountFollowingsFromUserID(userID int64) (int64, error) {
	// TODO: same as CountFollowersFromUserID
	count, err := q.count("relations", "following_id", "follower_id=$1", userID)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (q *PostgresQueries) DeleteRelation(followerID int64, followingID int64) error {
	if err := q.deleteWhere("relations", true, "follower_id=$1 AND following_id=$2", followerID, followingID); err != nil {
		return err // TODO: error handling when relation does not exist
	}

	return nil
}
