package repository

import "github.com/izruff/reviu-backend/internal/models"

func (q *PostgresQueries) CreateTag(newTag *models.Tag) (int64, error) {
	tagID, err := q.create("tags", []string{"tag", "hub"}, true, newTag)
	// TODO: error handling when form is incomplete or hub does not exist
	if err != nil {
		return 0, err
	}

	return tagID, nil
}

func (q *PostgresQueries) GetTagByID(id int64) (*models.Tag, error) {
	tag := &models.Tag{}
	if err := q.selectOne(tag, "tags", "*", "id=$1", id); err != nil {
		return nil, err // TODO: error handling when post does not exist
	}

	return tag, nil
}

// TODO: implement filters and preferences
func (q *PostgresQueries) GetTagsWithOptions(options interface{}) ([]*models.Tag, error) {
	return nil, nil
}
