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

func (q *PostgresQueries) DeleteTagFromPostWithID(postID int64, tagID int64) error {
	if err := q.deleteByPK("tagged_posts", strToAny{"post_id": postID, "tag_id": tagID}); err != nil {
		return err // TODO: error handling when user does not exist
	}

	return nil
}
