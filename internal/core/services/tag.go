package services

import (
	"github.com/izruff/reviu-backend/internal/core/domain"
	"gopkg.in/guregu/null.v3"
)

func (s *APIServices) CreateTag(tag string, hub string) (int64, *SvcError) {
	newTag := &domain.Tag{
		Tag: null.NewString(tag, true),
		Hub: null.NewString(hub, true),
	}

	tagID, err := s.repo.CreateTag(newTag)
	if err != nil {
		// TODO: error handling when topic already exists
		return 0, newErrInternal(err)
	}

	return tagID, nil
}

func (s *APIServices) GetTagByID(id int64) (*domain.Tag, *SvcError) {
	tag, err := s.repo.GetTagByID(id)
	if err != nil {
		return nil, newErrInternal(err) // TODO: error handling when tag does not exist
	}

	return tag, nil
}

func (s *APIServices) SearchTags(options *domain.SearchTagsOptions) ([]domain.Tag, *SvcError) {
	tags, err := s.repo.GetTagsWithOptions(options)
	if err != nil {
		return nil, newErrInternal(err) // TODO: error handling when there are incorrect options
	}

	return tags, nil
}
