package services

import (
	"github.com/izruff/reviu-backend/internal/core/domain"
	"gopkg.in/guregu/null.v3"
)

func (s *APIServices) CreateTag(tag string) (int64, *SvcError) {
	newTag := &domain.Tag{
		Tag: null.NewString(tag, true),
	}

	tagID, err := s.repo.CreateTag(newTag)
	if err != nil {
		// TODO: error handling when topic already exists
		return 0, newErrInternal(err)
	}

	return tagID, nil
}

func (s *APIServices) GetTagByID(id int64) (*domain.Tag, *SvcError) {
	tag, cErr := s.cache.GetTagFieldsByID(id, "tag", "createdAt")
	if cErr != nil {
		print(cErr.Err.Error()) // CacheTODO
	} else {
		return tag, nil
	}

	tag, err := s.repo.GetTagByID(id)
	if err != nil {
		return nil, newErrInternal(err) // TODO: error handling when tag does not exist
	}

	if cErr = s.cache.SetTagFieldsByID(id, tag); cErr != nil {
		print(cErr.Err.Error()) // CacheTODO
	}

	return tag, nil
}

func (s *APIServices) GetTagID(tag string) (int64, *SvcError) {
	// CacheTODO: Should we cache this? Do we even need this?
	topicID, err := s.repo.GetTagID(tag)
	if err != nil {
		return 0, newErrInternal(err) // TODO: error handling when topic does not exist
	}

	return topicID, nil
}

func (s *APIServices) SearchTags(options *domain.SearchTagsOptions) ([]domain.Tag, *SvcError) {
	tags, err := s.repo.GetTagsWithOptions(options)
	if err != nil {
		return nil, newErrInternal(err) // TODO: error handling when there are incorrect options
	}

	return tags, nil
}
