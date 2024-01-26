package repository

import (
	"errors"
	"strconv"

	"github.com/izruff/reviu-backend/internal/models"
)

func (q *PostgresQueries) CreateTopic(newTopic *models.Topic) (int64, error) {
	topicID, err := q.create("topics", []string{"topic", "hub"}, true, newTopic)
	// TODO: error handling when form is incomplete or hub does not exist
	if err != nil {
		return 0, err
	}

	return topicID, nil
}

func (q *PostgresQueries) GetTopicByID(id int64) (*models.Topic, error) {
	topic := &models.Topic{}
	if err := q.selectOne(topic, "topics", "*", "id=$1", id); err != nil {
		return nil, err // TODO: error handling when topic does not exist
	}

	return topic, nil
}

func (q *PostgresQueries) GetTopicID(topic string, hub string) (int64, error) {
	var topicID int64
	if err := q.selectOne(&topicID, "topics", "id", "topic=$1 AND hub=$2", topic, hub); err != nil {
		return 0, err // TODO: error handling when user does not exist
	}

	return topicID, nil
}

func (q *PostgresQueries) GetTopicsWithOptions(options *models.SearchTopicsOptions) ([]models.Topic, error) {
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

	var topics []models.Topic
	if err := q.selectAll(&topics, "topics", "*", whereQuery, orderBy, queryArgs...); err != nil {
		return nil, err
	}

	return topics, nil
}

func (q *PostgresQueries) UpdateTopicByID(id int64, description string) error {
	return nil // TODO: migrate the schema first
}
