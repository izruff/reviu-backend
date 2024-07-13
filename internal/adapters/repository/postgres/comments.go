package postgres

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/izruff/reviu-backend/internal/core/domain"
	"gopkg.in/guregu/null.v3"
)

func (r *PostgresRepository) CreateComment(newComment *domain.Comment) (int64, error) {
	// TODO: error handling when form is incomplete
	postID, err := r.create("comments", []string{"content", "author_id", "post_id", "parent_comment_id"}, true, newComment)
	if err != nil {
		return 0, err
	}

	return postID, nil
}

func (r *PostgresRepository) GetCommentByID(id int64) (*domain.Comment, error) {
	comment := &domain.Comment{}
	if err := r.selectOne(comment, "comments", "*", "id=$1", id); err != nil {
		return nil, err // TODO: error handling when comment does not exist
	}

	return comment, nil
}

func (r *PostgresRepository) GetCommentsWithOptions(options *domain.SearchCommentsOptions) ([]domain.Comment, error) {
	// TODO: error handling when post does not exist
	var whereQueries []string
	var orderBy string
	var queryArgs []interface{}
	argsIndex := 1

	if options.ParentCommentID.Valid {
		whereQueries = append(whereQueries, "parent_comment_id=$"+strconv.Itoa(argsIndex))
		queryArgs = append(queryArgs, options.ParentCommentID.Int64)
		argsIndex++
	} else if options.PostID.Valid {
		whereQueries = append(whereQueries, "post_id=$"+strconv.Itoa(argsIndex))
		queryArgs = append(queryArgs, options.PostID)
		argsIndex++
	}

	if options.Query != "" {
		whereQueries = append(whereQueries, "content %>> $"+strconv.Itoa(argsIndex))
		queryArgs = append(queryArgs, options.Query)
		argsIndex++
	}

	if options.SortBy == "similarity" {
		if options.Query == "" {
			return nil, errors.New("unexpected error: cannot sort by similarity when no query is given")
		}
		orderBy = "content <->>> $" + strconv.Itoa(argsIndex)
		queryArgs = append(queryArgs, options.Query)
		argsIndex++
	} else if options.SortBy == "popularity" {
		orderBy = "vote_count DESC, created_at DESC" // TODO: might implement a more complex popularity sorting
	} else if options.SortBy == "age-asc" || options.SortBy == "" {
		orderBy = "created_at DESC"
	} else if options.SortBy == "age-desc" {
		orderBy = "created_at ASC"
	} else {
		return nil, errors.New("unexpected error: invalid option for sort-by")
	}

	comments := []domain.Comment{}
	if err := r.selectAll(&comments, "comments", "*", strings.Join(whereQueries, " AND "), orderBy, queryArgs...); err != nil {
		return nil, err
	}

	return comments, nil
}

func (r *PostgresRepository) UpdateCommentByID(updatedComment *domain.Comment) error {
	if !updatedComment.ID.Valid {
		return errors.New("ID not provided")
	}

	var columns []string
	if updatedComment.Content.Valid {
		columns = append(columns, "content")
	}
	// TODO: error handling if nothing is updated

	updatedComment.UpdatedAt = null.NewTime(time.Now(), true)
	columns = append(columns, "updated_at")

	if err := r.updateByID("comments", columns, updatedComment); err != nil {
		return err // TODO: error handling when comment does not exist
	}

	return nil
}

func (r *PostgresRepository) MarkCommentAsDeletedByID(id int64, reason string, moderatorID int64) error {
	updatedComment := &domain.Comment{
		ID:                null.NewInt(id, true),
		ReasonForDeletion: null.NewString(reason, true),
		ModeratorID:       null.NewInt(moderatorID, true),
		DeletedAt:         null.NewTime(time.Now(), true),
	}
	if err := r.updateByID("comments", []string{"reason_for_deletion", "moderator_id", "deleted_at"}, updatedComment); err != nil {
		return err // TODO: error handling when comment does not exist
	}

	return nil
}
