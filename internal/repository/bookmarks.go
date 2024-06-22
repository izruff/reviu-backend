package repository

import "github.com/izruff/reviu-backend/internal/models"

func (q *PostgresQueries) CreateBookmark(newBookmark *models.Bookmark) error {
	// TODO: error handling when form is incomplete or already exists
	if _, err := q.create("bookmarks", []string{"post_id", "user_id"}, false, newBookmark); err != nil {
		return err
	}

	return nil
}

func (q *PostgresQueries) GetBookmarksFromUserID(userID int64) ([]models.Bookmark, error) {
	bookmarks := []models.Bookmark{}
	if err := q.selectAll(&bookmarks, "bookmarks", "post_id", "user_id=$1", "", userID); err != nil {
		return nil, err
	}

	return bookmarks, nil
}

func (q *PostgresQueries) CountBookmarksFromUserID(userID int64) (int64, error) {
	count, err := q.count("bookmarks", "post_id", "user_id=$1", userID)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (q *PostgresQueries) DeleteBookmark(postID int64, userID int64) error {
	if err := q.deleteWhere("bookmarks", true, "post_id=$1 AND user_id=$2", postID, userID); err != nil {
		return err // TODO: error handling when bookmark does not exist
	}

	return nil
}
