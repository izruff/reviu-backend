package repository

import "github.com/izruff/reviu-backend/internal/models"

func (q *PostgresQueries) CreateTopic(newTopic *models.Topic) (int64, error) {
	topicID, err := q.create("topics", []string{"topic", "hub"}, true, newTopic)
	// TODO: error handling when form is incomplete or hub does not exist
	if err != nil {
		return 0, err
	}

	return topicID, nil
}

func (q *PostgresQueries) GetTopicByID(id int64) (*models.Topic, error) {
	topic := &models.Topic{}
	if err := q.selectOne(topic, "topics", "*", "id=$1", id); err != nil {
		return nil, err // TODO: error handling when topic does not exist
	}

	return topic, nil
}

// TODO: implement filters and preferences
func (q *PostgresQueries) GetTopicsWithOptions(options strToAny) ([]*models.Topic, error) {
	return nil, nil
}
