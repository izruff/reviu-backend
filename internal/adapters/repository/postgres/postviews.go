package postgres

import "github.com/izruff/reviu-backend/internal/core/domain"

func (r *PostgresRepository) CreatePostView(newPostView *domain.PostView) error {
	// TODO: error handling when form is incomplete or already exists
	if _, err := r.create("post_views", []string{"post_id", "user_id"}, false, newPostView); err != nil {
		return err
	}

	return nil
}

func (r *PostgresRepository) GetPostViewValue(postID int64, userID int64) (bool, error) {
	var view domain.PostView
	if err := r.selectOne(&view, "post_views", "*", "post_id=$1 AND user_id=$2", postID, userID); err != nil {
		return false, nil // TODO: error handling for other internal errors
	}

	return true, nil
}

func (r *PostgresRepository) GetViewsFromPostID(postID int64) ([]domain.PostView, error) {
	views := []domain.PostView{}
	if err := r.selectAll(&views, "post_views", "*", "post_id=$1", "", postID); err != nil {
		return nil, err
	}

	return views, nil
}

func (r *PostgresRepository) CountViewsFromPostID(postID int64) (int64, error) {
	count, err := r.count("bookmarks", "user_id", "post_id=$1", postID)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *PostgresRepository) DeletePostView(postID int64, userID int64) error {
	if err := r.deleteWhere("post_views", true, "post_id=$1 AND user_id=$2", postID, userID); err != nil {
		return err // TODO: error handling when post view does not exist
	}

	return nil
}
