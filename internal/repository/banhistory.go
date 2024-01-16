package repository

import (
	"time"

	"github.com/izruff/reviu-backend/internal/models"
)

func (q *PostgresQueries) CreateBanHistory(newBanHistory *models.BanHistory) error {
	// TODO: error handling when form is incomplete or already exists
	if _, err := q.create("ban_history", []string{"start_time", "end_time", "reason", "user_id", "moderator_id"}, false, newBanHistory); err != nil {
		return err
	}

	return nil
}

func (q *PostgresQueries) GetBanHistoryFromUserID(userID int64) ([]*models.BanHistory, error) {
	return nil, nil // TODO
}

func (q *PostgresQueries) GetCurrentBanFromUserID(userID int64) (*models.BanHistory, error) {
	return nil, nil // TODO
}

func (q *PostgresQueries) DeleteBanHistory(startTime time.Time, userID int64) error {
	return nil // TODO
}
