package services

import (
	"github.com/izruff/reviu-backend/internal/models"
	"gopkg.in/guregu/null.v3"
)

func (s *APIServices) CreatePost(id int64) (int64, *SvcError) {
	newPost := &models.Post{}

	postID, err := s.queries.CreatePost(newPost)
	if err != nil {
		// TODO: error handling when post already exists
		return 0, newErrInternal(err)
	}

	return postID, nil
}

func (s *APIServices) GetPostByID(id int64) (*models.Post, *SvcError) {
	post, err := s.queries.GetPostByID(id)
	if err != nil {
		return nil, newErrInternal(err)
	}

	return post, nil
}

func (s *APIServices) UpdatePostByID(id int64, updatedPost *models.Post) *SvcError {
	// TODO: error handling when there are no changes
	updatedPost.ID.Int64 = id
	updatedPost.ID.Valid = true
	if err := s.queries.UpdatePostByID(updatedPost); err != nil {
		// TODO: error handling when post does not exist
		return newErrInternal(err)
	}

	return nil
}

func (s *APIServices) DeletePostByID(id int64, reasonForDeletion string, moderatorID int64) *SvcError {
	if err := s.queries.MarkPostAsDeletedByID(id, reasonForDeletion, moderatorID); err != nil {
		// TODO: error handling when post does not exist
		return newErrInternal(err)
	}

	return nil
}

func (s *APIServices) GetPostsByAuthorID(authorID int64) *SvcError {
	// posts, err := s.queries.GetPostsWithOptions()

	return nil // TODO
}

func (s *APIServices) BookmarkPostWithID(postID int64, userID int64) *SvcError {
	newBookmark := &models.Bookmark{
		PostID: null.NewInt(postID, true),
		UserID: null.NewInt(userID, true),
	}

	if err := s.queries.CreateBookmark(newBookmark); err != nil {
		// TODO: error handling when post or user does not exist
		return newErrInternal(err)
	}

	return nil
}
