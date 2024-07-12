package postgres

import "github.com/izruff/reviu-backend/internal/models"

func (q *PostgresQueries) CreatePostView(newPostView *models.PostView) error {
	// TODO: error handling when form is incomplete or already exists
	if _, err := q.create("post_views", []string{"post_id", "user_id"}, false, newPostView); err != nil {
		return err
	}

	return nil
}

func (q *PostgresQueries) GetPostViewValue(postID int64, userID int64) (bool, error) {
	var view models.PostView
	if err := q.selectOne(&view, "post_views", "*", "post_id=$1 AND user_id=$2", postID, userID); err != nil {
		return false, nil // TODO: error handling for other internal errors
	}

	return true, nil
}

func (q *PostgresQueries) GetViewsFromPostID(postID int64) ([]models.PostView, error) {
	views := []models.PostView{}
	if err := q.selectAll(&views, "post_views", "*", "post_id=$1", "", postID); err != nil {
		return nil, err
	}

	return views, nil
}

func (q *PostgresQueries) CountViewsFromPostID(postID int64) (int64, error) {
	count, err := q.count("bookmarks", "user_id", "post_id=$1", postID)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (q *PostgresQueries) DeletePostView(postID int64, userID int64) error {
	if err := q.deleteWhere("post_views", true, "post_id=$1 AND user_id=$2", postID, userID); err != nil {
		return err // TODO: error handling when post view does not exist
	}

	return nil
}
