package repository

import "github.com/izruff/reviu-backend/internal/models"

func (q *PostgresQueries) CreateSubscription(newSubscription *models.Subscription) error {
	// TODO: error handling when form is incomplete or already exists
	if _, err := q.create("subscriptions", []string{"topic_id", "user_id"}, false, newSubscription); err != nil {
		return err
	}

	return nil
}

func (q *PostgresQueries) GetSubscribersFromTopicID(topicID int64) ([]*models.Subscription, error) {
	var subscriptions []*models.Subscription
	if err := q.selectAll(subscriptions, "subscriptions", "user_id", "topic_id=$1", topicID); err != nil {
		return nil, err
	}

	return subscriptions, nil
}

func (q *PostgresQueries) CountSubscribersFromTopicID(topicID int64) (int64, error) {
	count, err := q.count("subscriptions", "user_id", "topic_id=$1", topicID)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (q *PostgresQueries) GetSubscribedTopicsFromUserID(userID int64) ([]*models.Subscription, error) {
	var subscriptions []*models.Subscription
	if err := q.selectAll(subscriptions, "subscriptions", "topic_id", "user_id=$1", userID); err != nil {
		return nil, err
	}

	return subscriptions, nil
}

func (q *PostgresQueries) CountSubscribedTopicsFromUserID(userID int64) (int64, error) {
	count, err := q.count("subscriptions", "topic_id", "user_id=$1", userID)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (q *PostgresQueries) DeleteSubscription(topicID int64, userID int64) error {
	if err := q.deleteWhere("subscriptions", true, "topic_id=$1 AND user_id=$2", topicID, userID); err != nil {
		return err // TODO: error handling when subscription does not exist
	}

	return nil
}
