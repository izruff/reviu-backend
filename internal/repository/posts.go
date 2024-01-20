package repository

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/izruff/reviu-backend/internal/models"
	"gopkg.in/guregu/null.v3"
)

func (q *PostgresQueries) CreatePost(newPost *models.Post) (int64, error) {
	postID, err := q.create("posts", []string{"title", "content", "author_id", "topic_id"}, true, newPost)
	// TODO: error handling when form is incomplete
	if err != nil {
		return 0, err
	}

	return postID, nil
}

func (q *PostgresQueries) GetPostByID(id int64) (*models.Post, error) {
	post := &models.Post{}
	if err := q.selectOne(post, "posts", "*", "id=$1", id); err != nil {
		return nil, err // TODO: error handling when post does not exist
	}

	return post, nil
}

func (q *PostgresQueries) GetPostsWithOptions(options *models.SearchPostsOptions) ([]models.Post, error) {
	var whereQueries []string
	var orderBy string
	var queryArgs []interface{}
	argsIndex := 1

	if options.Query == "" {
		return nil, errors.New("unexpected error: query is empty") // TODO: this should be allowed for browsing or searching recommendations
	}

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
		orderBy = "" // TODO
	} else if options.SortBy == "age-asc" || options.SortBy == "" {
		orderBy = "created_at DESC"
	} else if options.SortBy == "age-desc" {
		orderBy = "created_at ASC"
	} else {
		return nil, errors.New("unexpected error: invalid option for sort-by")
	}

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

	// TODO: handle topics and tags list

	var posts []models.Post
	if err := q.selectAll(&posts, "posts", "*", strings.Join(whereQueries, " AND "), orderBy, queryArgs...); err != nil {
		return nil, err
	}

	return posts, nil
}

func (q *PostgresQueries) UpdatePostByID(updatedPost *models.Post) error {
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

	if err := q.updateByID("posts", columns, updatedPost); err != nil {
		return err // TODO: error handling when post does not exist
	}

	return nil
}

func (q *PostgresQueries) MarkPostAsDeletedByID(id int64, reason string, moderatorID int64) error {
	updatedPost := &models.Post{
		ID:                null.NewInt(id, true),
		ReasonForDeletion: null.NewString(reason, true),
		ModeratorID:       null.NewInt(moderatorID, true),
		DeletedAt:         null.NewTime(time.Now(), true),
	}
	if err := q.updateByID("posts", []string{"reason_for_deletion", "moderator_id", "deleted_at"}, updatedPost); err != nil {
		return err // TODO: error handling when post does not exist
	}

	return nil
}
