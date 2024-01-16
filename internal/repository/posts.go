package repository

import (
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

// TODO: implement filters and preferences
func (q *PostgresQueries) GetPostsWithOptions(options interface{}) ([]*models.Post, error) {
	return nil, nil
}

func (q *PostgresQueries) UpdatePostByID(updatedPost *models.Post) error {
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
	columns = append(columns, "updatedAt")

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
