package postgres

import "github.com/izruff/reviu-backend/internal/core/domain"

func (r *PostgresRepository) CreateBookmark(newBookmark *domain.Bookmark) error {
	// TODO: error handling when form is incomplete or already exists
	if _, err := r.create("bookmarks", []string{"post_id", "user_id"}, false, newBookmark); err != nil {
		return err
	}

	return nil
}

func (r *PostgresRepository) GetBookmarksFromUserID(userID int64) ([]domain.Bookmark, error) {
	bookmarks := []domain.Bookmark{}
	if err := r.selectAll(&bookmarks, "bookmarks", "post_id", "user_id=$1", "", userID); err != nil {
		return nil, err
	}

	return bookmarks, nil
}

func (r *PostgresRepository) CountBookmarksFromUserID(userID int64) (int64, error) {
	count, err := r.count("bookmarks", "post_id", "user_id=$1", userID)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *PostgresRepository) DeleteBookmark(postID int64, userID int64) error {
	if err := r.deleteWhere("bookmarks", true, "post_id=$1 AND user_id=$2", postID, userID); err != nil {
		return err // TODO: error handling when bookmark does not exist
	}

	return nil
}
