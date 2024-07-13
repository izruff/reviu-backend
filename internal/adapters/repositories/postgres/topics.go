package postgres

import (
	"errors"
	"strconv"

	"github.com/izruff/reviu-backend/internal/core/domain"
)

func (r *PostgresRepository) CreateTopic(newTopic *domain.Topic) (int64, error) {
	topicID, err := r.create("topics", []string{"topic", "hub"}, true, newTopic)
	// TODO: error handling when form is incomplete or hub does not exist
	if err != nil {
		return 0, err
	}

	return topicID, nil
}

func (r *PostgresRepository) GetTopicByID(id int64) (*domain.Topic, error) {
	topic := &domain.Topic{}
	if err := r.selectOne(topic, "topics", "*", "id=$1", id); err != nil {
		return nil, err // TODO: error handling when topic does not exist
	}

	return topic, nil
}

func (r *PostgresRepository) GetTopicID(topic string, hub string) (int64, error) {
	var topicID int64
	if err := r.selectOne(&topicID, "topics", "id", "topic=$1 AND hub=$2", topic, hub); err != nil {
		return 0, err // TODO: error handling when user does not exist
	}

	return topicID, nil
}

func (r *PostgresRepository) GetTopicsWithOptions(options *domain.SearchTopicsOptions) ([]domain.Topic, error) {
	var whereQuery, orderBy string
	var queryArgs []interface{}
	argsIndex := 1

	if options.Query == "" {
		return nil, errors.New("unexpected error: query is empty") // TODO: this should be allowed for browsing
	}

	if options.SortBy == "similarity" || options.SortBy == "" {
		orderBy = "topic <-> $" + strconv.Itoa(argsIndex)
		queryArgs = append(queryArgs, options.Query)
		argsIndex++
	} else if options.SortBy == "popularity" {
		orderBy = "" // TODO
	} else { // default: similarity
		return nil, errors.New("unexpected error: invalid option for sort-by")
	}

	if options.MustMatch == "left" {
		whereQuery = "topic ILIKE $" + strconv.Itoa(argsIndex)
		queryArgs = append(queryArgs, options.Query+"%")
		argsIndex++
	} else if options.MustMatch == "substring" {
		whereQuery = "topic ILIKE $" + strconv.Itoa(argsIndex)
		queryArgs = append(queryArgs, "%"+options.Query+"%")
		argsIndex++
	} else if options.MustMatch == "none" || options.MustMatch == "" {
		whereQuery = "topic % $" + strconv.Itoa(argsIndex)
		queryArgs = append(queryArgs, options.Query)
		argsIndex++
	} else {
		return nil, errors.New("unexpected error: invalid option for must-match")
	}

	topics := []domain.Topic{}
	if err := r.selectAll(&topics, "topics", "*", whereQuery, orderBy, queryArgs...); err != nil {
		return nil, err
	}

	return topics, nil
}

func (r *PostgresRepository) UpdateTopicByID(id int64, description string) error {
	return nil // TODO: migrate the schema first
}
