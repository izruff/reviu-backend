package postgres

import "github.com/izruff/reviu-backend/internal/core/domain"

func (r *PostgresRepository) CreateSubscription(newSubscription *domain.Subscription) error {
	// TODO: error handling when form is incomplete or already exists
	if _, err := r.create("subscriptions", []string{"topic_id", "user_id"}, false, newSubscription); err != nil {
		return err
	}

	return nil
}

func (r *PostgresRepository) GetSubscribersFromTopicID(topicID int64) ([]domain.Subscription, error) {
	subscriptions := []domain.Subscription{}
	if err := r.selectAll(&subscriptions, "subscriptions", "user_id", "topic_id=$1", "", topicID); err != nil {
		return nil, err
	}

	return subscriptions, nil
}

func (r *PostgresRepository) CountSubscribersFromTopicID(topicID int64) (int64, error) {
	count, err := r.count("subscriptions", "user_id", "topic_id=$1", topicID)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *PostgresRepository) GetSubscribedTopicsFromUserID(userID int64) ([]domain.Subscription, error) {
	subscriptions := []domain.Subscription{}
	if err := r.selectAll(&subscriptions, "subscriptions", "topic_id", "user_id=$1", "", userID); err != nil {
		return nil, err
	}

	return subscriptions, nil
}

func (r *PostgresRepository) CountSubscribedTopicsFromUserID(userID int64) (int64, error) {
	count, err := r.count("subscriptions", "topic_id", "user_id=$1", userID)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *PostgresRepository) DeleteSubscription(topicID int64, userID int64) error {
	if err := r.deleteWhere("subscriptions", true, "topic_id=$1 AND user_id=$2", topicID, userID); err != nil {
		return err // TODO: error handling when subscription does not exist
	}

	return nil
}
