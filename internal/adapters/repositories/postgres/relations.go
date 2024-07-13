package postgres

import "github.com/izruff/reviu-backend/internal/core/domain"

func (r *PostgresRepository) CreateRelation(newRelation *domain.Relation) error {
	// TODO: error handling when form is incomplete or already exists
	if _, err := r.create("relations", []string{"follower_id", "following_id"}, false, newRelation); err != nil {
		return err
	}

	return nil
}

func (r *PostgresRepository) GetFollowersFromUserID(userID int64) ([]domain.Relation, error) {
	relations := []domain.Relation{}
	if err := r.selectAll(&relations, "relations", "follower_id", "following_id=$1", "", userID); err != nil {
		return nil, err
	}

	return relations, nil
}

func (r *PostgresRepository) CountFollowersFromUserID(userID int64) (int64, error) {
	// TODO: possibly optimize this by creating a new column in the database
	count, err := r.count("relations", "follower_id", "following_id=$1", userID)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *PostgresRepository) GetFollowingsFromUserID(userID int64) ([]domain.Relation, error) {
	relations := []domain.Relation{}
	if err := r.selectAll(&relations, "relations", "following_id", "follower_id=$1", "", userID); err != nil {
		return nil, err
	}

	return relations, nil
}

func (r *PostgresRepository) CountFollowingsFromUserID(userID int64) (int64, error) {
	// TODO: same as CountFollowersFromUserID
	count, err := r.count("relations", "following_id", "follower_id=$1", userID)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *PostgresRepository) DeleteRelation(followerID int64, followingID int64) error {
	if err := r.deleteWhere("relations", true, "follower_id=$1 AND following_id=$2", followerID, followingID); err != nil {
		return err // TODO: error handling when relation does not exist
	}

	return nil
}
