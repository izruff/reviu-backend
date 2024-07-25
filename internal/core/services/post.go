package services

import (
	"time"

	"github.com/izruff/reviu-backend/internal/core/domain"
	"gopkg.in/guregu/null.v3"
)

func (s *APIServices) CreatePost(title string, content string, authorID int64, topic string, tags []string) (int64, *SvcError) {
	newTopic := &domain.Topic{
		Topic: null.NewString(topic, true),
	}
	topicID, err := s.repo.CreateTopic(newTopic)
	if err != nil {
		if true { // TODO: error handling when tag already exists (replace true with err != ...)
			topicID, err = s.repo.GetTopicID(topic)
			if err != nil {
				return 0, newErrInternal(err)
			}
		} else {
			return 0, newErrInternal(err)
		}
	}

	newPost := &domain.Post{
		Title:    null.NewString(title, true),
		Content:  null.NewString(content, true),
		AuthorID: null.NewInt(authorID, true),
		TopicID:  null.NewInt(topicID, true),
	}

	postID, err := s.repo.CreatePost(newPost)
	if err != nil {
		// TODO: error handling when post already exists
		return 0, newErrInternal(err)
	}

	for _, tag := range tags {
		newTag := &domain.Tag{
			Tag: null.NewString(tag, true),
		}
		tagID, err := s.repo.CreateTag(newTag) // return tagID or not?
		if err != nil && true {                // TODO: error handling when tag already exists (replace true with err != ...)
			return 0, newErrInternal(err)
		}

		newTaggedPost := &domain.TaggedPost{
			PostID: null.NewInt(postID, true),
			TagID:  null.NewInt(tagID, true),
		}
		if err := s.repo.CreateTaggedPost(newTaggedPost); err != nil {
			return 0, newErrInternal(err)
		}
	}

	return postID, nil
}

func (s *APIServices) GetPostByID(id int64) (*domain.Post, *SvcError) {
	post, cErr := s.cache.GetPostFieldsByID(id, "title", "content", "authorId", "topicId", "createdAt", "updatedAt", "deletedDetails", "viewCount", "voteCount")
	if cErr != nil {
		print(cErr.Err.Error()) // CacheTODO
	} else {
		return post, nil
	}

	post, err := s.repo.GetPostByID(id)
	if err != nil {
		// TODO: error handling when post does not exist
		return nil, newErrInternal(err)
	}

	if cErr = s.cache.SetPostFieldsByID(id, post); cErr != nil {
		print(cErr.Err.Error()) // CacheTODO
	}

	return post, nil
}

func (s *APIServices) GetPostInteractionsByUserID(id int64, userID int64) (bool, *null.Bool, *SvcError) {
	// CacheTODO: Should we cache this?
	viewed, err := s.repo.GetPostViewValue(id, userID)
	if err != nil {
		return false, nil, newErrInternal(err)
	}

	voted, err := s.repo.GetPostVoteValue(id, userID)
	if err != nil {
		return false, nil, newErrInternal(err)
	}

	return viewed, voted, nil
}

func (s *APIServices) UpdatePostByID(id int64, updatedPost *domain.Post) *SvcError {
	updatedPost.ID.SetValid(id)
	if err := s.repo.UpdatePostByID(updatedPost); err != nil {
		// TODO: error handling when post does not exist
		return newErrInternal(err)
	}

	if cErr := s.cache.SetPostFieldsByID(id, updatedPost); cErr != nil {
		print(cErr.Err.Error()) // CacheTODO
	}

	return nil
}

func (s *APIServices) MarkPostAsDeletedByID(id int64, deletedAt time.Time, reasonForDeletion string, moderatorID int64) *SvcError {
	if err := s.repo.MarkPostAsDeletedByID(id, deletedAt, reasonForDeletion, moderatorID); err != nil {
		// TODO: error handling when post does not exist
		return newErrInternal(err)
	}

	// cache must update successfully
	updatedPost := &domain.Post{
		DeletedAt:         null.NewTime(deletedAt, true),
		ReasonForDeletion: null.NewString(reasonForDeletion, true),
		ModeratorID:       null.NewInt(moderatorID, true),
	}
	if cErr := s.cache.SetPostFieldsByID(id, updatedPost); cErr != nil {
		return newErrInternal(cErr.Err) // CacheTODO
	}

	return nil
}

func (s *APIServices) ViewPost(id int64, userID int64) *SvcError {
	// CacheTODO: Should update view count and make sure repository function returns delta value
	newView := &domain.PostView{
		PostID: null.NewInt(id, true),
		UserID: null.NewInt(userID, true),
	}

	if err := s.repo.CreatePostView(newView); err != nil {
		return newErrInternal(err) // TODO: error handling when post view already exists
	}

	return nil
}

func (s *APIServices) VotePost(id int64, userID int64, up null.Bool) *SvcError {
	// TODO: this logic is assuming there is no possibility for other weird internal errors
	// CacheTODO: Should update vote count and make sure repository function returns delta value
	if !up.Valid {
		s.repo.DeletePostVote(id, userID)
		return nil
	}
	if err := s.repo.UpdatePostVote(up.Bool, id, userID); err != nil {
		newVote := &domain.PostVote{
			Up:     null.NewBool(up.Bool, true),
			PostID: null.NewInt(id, true),
			UserID: null.NewInt(userID, true),
		}
		if err := s.repo.CreatePostVote(newVote); err != nil {
			return newErrInternal(err)
		}
	}

	return nil
}

func (s *APIServices) SearchPosts(options *domain.SearchPostsOptions) ([]domain.Post, *SvcError) {
	posts, err := s.repo.GetPostsWithOptions(options)
	if err != nil {
		return nil, newErrInternal(err) // TODO: error handling when there are incorrect options
	}

	return posts, nil
}

func (s *APIServices) GetPostsByAuthorID(authorID int64) *SvcError {
	// posts, err := s.repo.GetPostsWithOptions()

	return nil // TODO
}

func (s *APIServices) BookmarkPostWithID(postID int64, userID int64) *SvcError {
	newBookmark := &domain.Bookmark{
		PostID: null.NewInt(postID, true),
		UserID: null.NewInt(userID, true),
	}

	if err := s.repo.CreateBookmark(newBookmark); err != nil {
		// TODO: error handling when post or user does not exist
		return newErrInternal(err)
	}

	return nil
}
