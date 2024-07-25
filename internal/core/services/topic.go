package services

import (
	"github.com/izruff/reviu-backend/internal/core/domain"
	"gopkg.in/guregu/null.v3"
)

func (s *APIServices) CreateTopic(topic string) (int64, *SvcError) {
	newTopic := &domain.Topic{
		Topic: null.NewString(topic, true),
	}

	topicID, err := s.repo.CreateTopic(newTopic)
	if err != nil {
		// TODO: error handling when topic already exists
		return 0, newErrInternal(err)
	}

	return topicID, nil
}

func (s *APIServices) GetTopicByID(id int64) (*domain.Topic, *SvcError) {
	topic, err := s.repo.GetTopicByID(id)
	if err != nil {
		return nil, newErrInternal(err) // TODO: error handling when topic does not exist
	}

	return topic, nil
}

func (s *APIServices) GetTopicID(topic string) (int64, *SvcError) {
	topicID, err := s.repo.GetTopicID(topic)
	if err != nil {
		return 0, newErrInternal(err) // TODO: error handling when topic does not exist
	}

	return topicID, nil
}

func (s *APIServices) UpdateTopicByID(id int64, description string) *SvcError {
	if err := s.repo.UpdateTopicByID(id, description); err != nil {
		// TODO: error handling when user does not exist
		return newErrInternal(err)
	}

	return nil
}

func (s *APIServices) SearchTopics(options *domain.SearchTopicsOptions) ([]domain.Topic, *SvcError) {
	return nil, nil // TODO
}
