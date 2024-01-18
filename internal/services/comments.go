package services

import (
	"github.com/izruff/reviu-backend/internal/models"
	"gopkg.in/guregu/null.v3"
)

func (s *APIServices) CreateComment(content string, authorID int64, postID int64, parentCommentID null.Int) (int64, *SvcError) {
	newComment := &models.Comment{
		Content:         null.NewString(content, true),
		AuthorID:        null.NewInt(authorID, true),
		PostID:          null.NewInt(postID, true),
		ParentCommentID: parentCommentID,
	}

	commentID, err := s.queries.CreateComment(newComment)
	if err != nil {
		// TODO: error handling when post already exists
		return 0, newErrInternal(err)
	}

	return commentID, nil
}

func (s *APIServices) GetCommentByID(commentID int64, postID int64) (*models.Comment, *SvcError) {
	return nil, nil // TODO
}

func (s *APIServices) UpdateCommentByID(commentID int64, postID int64, comment string) *SvcError {
	return nil // TODO
}

func (s *APIServices) DeleteCommentByID(commentID int64, postID int64, reasonForDeletion string, moderatorID int64) *SvcError {
	return nil // TODO
}

func (s *APIServices) SearchCommentsInPost(options *models.SearchCommentsOptions) ([]*models.Comment, *SvcError) {
	return nil, nil // TODO
}
