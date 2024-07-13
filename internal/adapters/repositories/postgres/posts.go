package postgres

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/izruff/reviu-backend/internal/core/domain"
	"gopkg.in/guregu/null.v3"
)

func (r *PostgresRepository) CreatePost(newPost *domain.Post) (int64, error) {
	postID, err := r.create("posts", []string{"title", "content", "author_id", "topic_id"}, true, newPost)
	// TODO: error handling when form is incomplete
	if err != nil {
		return 0, err
	}

	return postID, nil
}

func (r *PostgresRepository) GetPostByID(id int64) (*domain.Post, error) {
	post := &domain.Post{}
	if err := r.selectOne(post, "posts", "*", "id=$1", id); err != nil {
		return nil, err // TODO: error handling when post does not exist
	}

	return post, nil
}

func (r *PostgresRepository) GetPostsWithOptions(options *domain.SearchPostsOptions) ([]domain.Post, error) {
	var whereQueries []string
	var orderBy string
	var queryArgs []interface{}
	argsIndex := 1

	if options.SortBy == "similarity" {
		if options.MatchWith == "title" {
			orderBy = "title <->>> $" + strconv.Itoa(argsIndex)
			queryArgs = append(queryArgs, options.Query)
			argsIndex++
		} else {
			orderBy = "GREATEST(title <->>> $" + strconv.Itoa(argsIndex) + ", content <->>> $" + strconv.Itoa(argsIndex+1) + ")"
			queryArgs = append(queryArgs, options.Query, options.Query)
			argsIndex += 2
		}
	} else if options.SortBy == "popularity" {
		orderBy = "vote_count DESC, created_at DESC" // TODO: might implement a more complex popularity sorting
	} else if options.SortBy == "age-asc" || options.SortBy == "" {
		orderBy = "created_at DESC"
	} else if options.SortBy == "age-desc" {
		orderBy = "created_at ASC"
	} else {
		return nil, errors.New("unexpected error: invalid option for sort-by")
	}

	if options.Query != "" {
		if options.MatchWith == "title" {
			whereQueries = append(whereQueries, "title %>> $"+strconv.Itoa(argsIndex))
			queryArgs = append(queryArgs, options.Query)
			argsIndex++
		} else if options.MatchWith == "all" || options.MatchWith == "" {
			whereQueries = append(whereQueries, "(title %>> $"+strconv.Itoa(argsIndex)+" OR content %>> $"+strconv.Itoa(argsIndex+1)+")")
			queryArgs = append(queryArgs, options.Query, options.Query)
			argsIndex += 2
		} else {
			return nil, errors.New("unexpected error: invalid option for match-with")
		}
	}

	if len(options.Authors) > 0 {
		subsubquery := ""
		for i, author := range options.Authors {
			if i > 0 {
				subsubquery += ","
			}
			subsubquery += "$" + strconv.Itoa(argsIndex)
			queryArgs = append(queryArgs, author)
			argsIndex++
		}
		whereQueries = append(whereQueries, "author_id IN (SELECT id FROM users WHERE username IN ("+subsubquery+"))")
	}

	// TODO: handle topics and tags list (must include both name and hub since names are not necessarily distinct)

	posts := []domain.Post{}
	if err := r.selectAll(&posts, "posts", "*", strings.Join(whereQueries, " AND "), orderBy, queryArgs...); err != nil {
		return nil, err
	}

	return posts, nil
}

func (r *PostgresRepository) CountPostsFromAuthorID(userID int64) (int64, error) {
	// TODO: possibly optimize this by creating a new column in the database
	// TODO: handle logic for deleted and/or edited posts
	count, err := r.count("posts", "id", "author_id=$1", userID)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *PostgresRepository) UpdatePostByID(updatedPost *domain.Post) error {
	if !updatedPost.ID.Valid {
		return errors.New("ID not provided")
	}

	var columns []string
	if updatedPost.Title.Valid {
		columns = append(columns, "title")
	}
	if updatedPost.Content.Valid {
		columns = append(columns, "content")
	}
	if updatedPost.TopicID.Valid {
		columns = append(columns, "topic_id")
	}
	// TODO: error handling if nothing is updated

	updatedPost.UpdatedAt = null.NewTime(time.Now(), true)
	columns = append(columns, "updated_at")

	if err := r.updateByID("posts", columns, updatedPost); err != nil {
		return err // TODO: error handling when post does not exist
	}

	return nil
}

func (r *PostgresRepository) MarkPostAsDeletedByID(id int64, reason string, moderatorID int64) error {
	updatedPost := &domain.Post{
		ID:                null.NewInt(id, true),
		ReasonForDeletion: null.NewString(reason, true),
		ModeratorID:       null.NewInt(moderatorID, true),
		DeletedAt:         null.NewTime(time.Now(), true),
	}
	if err := r.updateByID("posts", []string{"reason_for_deletion", "moderator_id", "deleted_at"}, updatedPost); err != nil {
		return err // TODO: error handling when post does not exist
	}

	return nil
}
