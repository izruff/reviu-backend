package postgres

import (
	"errors"
	"strconv"

	"github.com/izruff/reviu-backend/internal/core/domain"
)

func (r *PostgresRepository) CreateTag(newTag *domain.Tag) (int64, error) {
	tagID, err := r.create("tags", []string{"tag", "hub"}, true, newTag)
	// TODO: error handling when form is incomplete or hub does not exist
	if err != nil {
		return 0, err
	}

	return tagID, nil
}

func (r *PostgresRepository) GetTagByID(id int64) (*domain.Tag, error) {
	tag := &domain.Tag{}
	if err := r.selectOne(tag, "tags", "*", "id=$1", id); err != nil {
		return nil, err // TODO: error handling when tag does not exist
	}

	return tag, nil
}

func (r *PostgresRepository) GetTagsWithOptions(options *domain.SearchTagsOptions) ([]domain.Tag, error) {
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

	tags := []domain.Tag{}
	if err := r.selectAll(&tags, "tags", "*", whereQuery, orderBy, queryArgs...); err != nil {
		return nil, err
	}

	return tags, nil
}
