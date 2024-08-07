package services

import (
	"github.com/izruff/reviu-backend/internal/core/domain"
	"gopkg.in/guregu/null.v3"
)

func (s *APIServices) CreateComment(content string, authorID int64, postID null.Int, parentCommentID null.Int) (int64, *SvcError) {
	if !postID.Valid {
		if !parentCommentID.Valid {
			return 0, newErrInvalidUserInput([]string{"postID", "parentCommentID"}) // TODO: make a new type of error for this
		}
		parentComment, err := s.repo.GetCommentByID(parentCommentID.Int64)
		if err != nil {
			return 0, newErrInternal(err)
		}
		postID.Int64 = parentComment.PostID.Int64
		postID.Valid = true
	}

	newComment := &domain.Comment{
		Content:         null.NewString(content, true),
		AuthorID:        null.NewInt(authorID, true),
		PostID:          postID,
		ParentCommentID: parentCommentID,
	}

	commentID, err := s.repo.CreateComment(newComment)
	if err != nil {
		// TODO: error handling when post already exists
		return 0, newErrInternal(err)
	}

	return commentID, nil
}

func (s *APIServices) GetCommentByID(commentID int64) (*domain.Comment, *SvcError) {
	comment, err := s.repo.GetCommentByID(commentID)
	if err != nil {
		// TODO: error handling when comment does not exist
		return nil, newErrInternal(err)
	}

	return comment, nil
}

func (s *APIServices) UpdateCommentByID(commentID int64, content string) *SvcError {
	updatedComment := &domain.Comment{
		ID:      null.NewInt(commentID, true),
		Content: null.NewString(content, true),
	}
	if err := s.repo.UpdateCommentByID(updatedComment); err != nil {
		// TODO: error handling when comment does not exist
		return newErrInternal(err)
	}

	return nil
}

func (s *APIServices) MarkCommentAsDeletedByID(commentID int64, postID int64, reasonForDeletion string, moderatorID int64) *SvcError {
	if err := s.repo.MarkCommentAsDeletedByID(commentID, reasonForDeletion, moderatorID); err != nil {
		// TODO: error handling when comment does not exist
		return newErrInternal(err)
	}

	return nil
}

func (s *APIServices) VoteComment(id int64, userID int64, up null.Bool) *SvcError {
	// TODO: this logic is assuming there is no possibility for other weird internal errors
	if !up.Valid {
		s.repo.DeleteCommentVote(id, userID)
		return nil
	}
	if err := s.repo.UpdateCommentVote(up.Bool, id, userID); err != nil {
		newVote := &domain.CommentVote{
			Up:        null.NewBool(up.Bool, true),
			CommentID: null.NewInt(id, true),
			UserID:    null.NewInt(userID, true),
		}
		if err := s.repo.CreateCommentVote(newVote); err != nil {
			return newErrInternal(err)
		}
	}

	return nil
}

func (s *APIServices) SearchComments(options *domain.SearchCommentsOptions) ([]domain.Comment, *SvcError) {
	comments, err := s.repo.GetCommentsWithOptions(options)
	if err != nil {
		return nil, newErrInternal(err) // TODO: error handling when there are incorrect options
	}

	return comments, nil
}
