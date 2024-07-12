package services

import (
	"github.com/izruff/reviu-backend/internal/models"
	"gopkg.in/guregu/null.v3"
)

func (s *APIServices) CreateTopic(topic string, hub string) (int64, *SvcError) {
	newTopic := &models.Topic{
		Topic: null.NewString(topic, true),
		Hub:   null.NewString(hub, true),
	}

	topicID, err := s.queries.CreateTopic(newTopic)
	if err != nil {
		// TODO: error handling when topic already exists
		return 0, newErrInternal(err)
	}

	return topicID, nil
}

func (s *APIServices) GetTopicByID(id int64) (*models.Topic, *SvcError) {
	topic, err := s.queries.GetTopicByID(id)
	if err != nil {
		return nil, newErrInternal(err) // TODO: error handling when topic does not exist
	}

	return topic, nil
}

func (s *APIServices) UpdateTopicByID(id int64, description string) *SvcError {
	if err := s.queries.UpdateTopicByID(id, description); err != nil {
		// TODO: error handling when user does not exist
		return newErrInternal(err)
	}

	return nil
}

func (s *APIServices) SearchTopics(options *models.SearchTopicsOptions) ([]models.Topic, *SvcError) {
	return nil, nil // TODO
}
