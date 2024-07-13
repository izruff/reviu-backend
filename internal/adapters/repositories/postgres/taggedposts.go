package postgres

import "github.com/izruff/reviu-backend/internal/core/domain"

func (r *PostgresRepository) CreateTaggedPost(newTaggedPost *domain.TaggedPost) error {
	// TODO: error handling when form is incomplete or already exists
	if _, err := r.create("tagged_posts", []string{"post_id", "tag_id"}, false, newTaggedPost); err != nil {
		return err
	}

	return nil
}

func (r *PostgresRepository) GetTagsFromPostID(postID int64) ([]domain.Tag, error) {
	return nil, nil // TODO: maybe just merge with GetTagsWithOptions
}

// TODO: might implement a function which returns a (Postgres) view of the tags in a post/vice versa

func (r *PostgresRepository) DeleteTagFromPost(postID int64, tagID int64) error {
	if err := r.deleteWhere("tagged_posts", true, "post_id=$1 AND tag_id=$2", postID, tagID); err != nil {
		return err // TODO: error handling when tagged post does not exist
	}

	return nil
}
