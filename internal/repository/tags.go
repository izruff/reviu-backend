package repository

import (
	"errors"
	"strconv"

	"github.com/izruff/reviu-backend/internal/models"
)

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
		return nil, err // TODO: error handling when tag does not exist
	}

	return tag, nil
}

func (q *PostgresQueries) GetTagsWithOptions(options *models.SearchTagsOptions) ([]models.Tag, error) {
	var whereQuery, orderBy string
	var queryArgs []interface{}
	argsIndex := 1

	if options.Query == "" {
		return nil, errors.New("unexpected error: query is empty")
	}

	if options.SortBy == "similarity" {
		orderBy = "tag <-> $" + strconv.Itoa(argsIndex)
		queryArgs = append(queryArgs, options.Query)
		argsIndex++
	} else if options.SortBy == "popularity" {
		orderBy = "" // TODO
	} else {
		return nil, errors.New("unexpected error: invalid option for sort-by")
	}

	if options.MustMatch == "left" {
		whereQuery = "tag ILIKE $" + strconv.Itoa(argsIndex)
		queryArgs = append(queryArgs, options.Query+"%")
		argsIndex++
	} else if options.MustMatch == "substring" {
		whereQuery = "tag ILIKE $" + strconv.Itoa(argsIndex)
		queryArgs = append(queryArgs, "%"+options.Query+"%")
		argsIndex++
	} else if options.MustMatch == "none" || options.MustMatch == "" {
		whereQuery = "tag % $" + strconv.Itoa(argsIndex)
		queryArgs = append(queryArgs, options.Query)
		argsIndex++
	} else {
		return nil, errors.New("unexpected error: invalid option for must-match")
	}

	tags := []models.Tag{}
	if err := q.selectAll(&tags, "tags", "*", whereQuery, orderBy, queryArgs...); err != nil {
		return nil, err
	}

	return tags, nil
}
