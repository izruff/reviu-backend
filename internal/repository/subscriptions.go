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
	return nil, nil // TODO
}

func (q *PostgresQueries) CountSubscribersFromTopicID(topicID int64) (int64, error) {
	return 0, nil // TODO
}

func (q *PostgresQueries) GetSubscribedTopicsFromUserID(userID int64) ([]*models.Subscription, error) {
	return nil, nil // TODO
}

func (q *PostgresQueries) CountSubscribedTopicsFromUserID(userID int64) (int64, error) {
	return 0, nil // TODO
}

func (q *PostgresQueries) DeleteSubscription(topicID int64, userID int64) error {
	return nil // TODO
}
