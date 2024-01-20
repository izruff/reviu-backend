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

// Initially, I intended for the ID to start at 1 for each distinct postID, hence why the three
// functions below needs two arguments, but I did not find a simple way to do this. So instead, here
// the ID is the primary key and always increments by 1, so postID argument is not actually necessary.

func (s *APIServices) GetCommentByID(commentID int64, postID int64) (*models.Comment, *SvcError) {
	comment, err := s.queries.GetCommentByID(commentID)
	if err != nil {
		// TODO: error handling when comment does not exist
		return nil, newErrInternal(err)
	}

	return comment, nil
}

func (s *APIServices) UpdateCommentByID(commentID int64, postID int64, content string) *SvcError {
	updatedComment := &models.Comment{
		ID:      null.NewInt(commentID, true),
		Content: null.NewString(content, true),
	}
	if err := s.queries.UpdateCommentByID(updatedComment); err != nil {
		// TODO: error handling when comment does not exist
		return newErrInternal(err)
	}

	return nil
}

func (s *APIServices) MarkCommentAsDeletedByID(commentID int64, postID int64, reasonForDeletion string, moderatorID int64) *SvcError {
	if err := s.queries.MarkCommentAsDeletedByID(commentID, reasonForDeletion, moderatorID); err != nil {
		// TODO: error handling when comment does not exist
		return newErrInternal(err)
	}

	return nil
}

func (s *APIServices) SearchCommentsInPost(options *models.SearchCommentsOptions) ([]models.Comment, *SvcError) {
	return nil, nil // TODO
}
