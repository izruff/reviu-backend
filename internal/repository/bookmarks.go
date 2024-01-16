package repository

import "github.com/izruff/reviu-backend/internal/models"

func (q *PostgresQueries) CreateBookmark(newBookmark *models.Bookmark) error {
	// TODO: error handling when form is incomplete or already exists
	if _, err := q.create("bookmarks", []string{"post_id", "user_id"}, false, newBookmark); err != nil {
		return err
	}

	return nil
}

func (q *PostgresQueries) GetBookmarksFromUserID(userID int64) ([]*models.Bookmark, error) {
	return nil, nil // TODO
}

func (q *PostgresQueries) CountBookmarksFromUserID(userID int64) (int64, error) {
	return 0, nil // TODO
}

func (q *PostgresQueries) DeleteBookmark(postID int64, userID int64) error {
	return nil // TODO
}
