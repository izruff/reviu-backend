package services

import (
	"github.com/izruff/reviu-backend/internal/models"
	"gopkg.in/guregu/null.v3"
)

func (s *APIServices) CreatePost(title string, content string, authorID int64, topicID int64, tags []string) (int64, *SvcError) {
	newPost := &models.Post{
		Title:    null.NewString(title, true),
		Content:  null.NewString(content, true),
		AuthorID: null.NewInt(authorID, true),
		TopicID:  null.NewInt(topicID, true),
	}

	postID, err := s.queries.CreatePost(newPost)
	if err != nil {
		// TODO: error handling when post already exists
		return 0, newErrInternal(err)
	}

	topic, _ := s.queries.GetTopicByID(topicID)
	hub := topic.Hub
	for _, tag := range tags {
		newTag := &models.Tag{
			Tag: null.NewString(tag, true),
			Hub: hub,
		}
		tagID, err := s.queries.CreateTag(newTag) // return tagID or not?
		if err != nil && true {                   // TODO: error handling when tag already exists (replace true with condition)
			return 0, newErrInternal(err)
		}

		newTaggedPost := &models.TaggedPost{
			PostID: null.NewInt(postID, true),
			TagID:  null.NewInt(tagID, true),
		}
		if err := s.queries.CreateTaggedPost(newTaggedPost); err != nil {
			return 0, newErrInternal(err)
		}
	}

	return postID, nil
}

func (s *APIServices) GetPostByID(id int64) (*models.Post, *SvcError) {
	post, err := s.queries.GetPostByID(id)
	if err != nil {
		// TODO: error handling when post does not exist
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

func (s *APIServices) MarkPostAsDeletedByID(id int64, reasonForDeletion string, moderatorID int64) *SvcError {
	if err := s.queries.MarkPostAsDeletedByID(id, reasonForDeletion, moderatorID); err != nil {
		// TODO: error handling when post does not exist
		return newErrInternal(err)
	}

	return nil
}

func (s *APIServices) VotePost(id int64, userID int64, up bool) *SvcError {
	return nil // TODO
}

func (s *APIServices) SearchPosts(options *models.SearchPostsOptions) ([]models.Post, *SvcError) {
	users, err := s.queries.GetPostsWithOptions(options)
	if err != nil {
		return nil, newErrInternal(err) // TODO: error handling when there are incorrect options
	}

	return users, nil // TODO
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
