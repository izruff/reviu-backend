package repository

import "github.com/izruff/reviu-backend/internal/models"

func (q *PostgresQueries) CreateTaggedPost(newTaggedPost *models.TaggedPost) error {
	// TODO: error handling when form is incomplete or already exists
	if _, err := q.create("tagged_posts", []string{"post_id", "tag_id"}, false, newTaggedPost); err != nil {
		return err
	}

	return nil
}

func (q *PostgresQueries) GetTagsFromPostID(postID int64) ([]*models.Tag, error) {
	return nil, nil // TODO
}

// TODO: might implement a function which returns a (Postgres) view of the tags in a post/vice versa

func (q *PostgresQueries) DeleteTagFromPost(postID int64, tagID int64) error {
	if err := q.deleteWhere("tagged_posts", true, "post_id=$1 AND tag_id=$2", postID, tagID); err != nil {
		return err // TODO: error handling when tagged post does not exist
	}

	return nil
}
