package services

import (
	"github.com/izruff/reviu-backend/internal/models"
	"gopkg.in/guregu/null.v3"
)

func (s *APIServices) CreateTag(tag string, hub string) (int64, *SvcError) {
	newTag := &models.Tag{
		Tag: null.NewString(tag, true),
		Hub: null.NewString(hub, true),
	}

	tagID, err := s.queries.CreateTag(newTag)
	if err != nil {
		// TODO: error handling when topic already exists
		return 0, newErrInternal(err)
	}

	return tagID, nil
}

func (s *APIServices) GetTagByID(id int64) (*models.Tag, *SvcError) {
	tag, err := s.queries.GetTagByID(id)
	if err != nil {
		return nil, newErrInternal(err) // TODO: error handling when tag does not exist
	}

	return tag, nil
}

func (s *APIServices) SearchTags(options *models.SearchTagsOptions) ([]*models.Tag, *SvcError) {
	return nil, nil // TODO
}
