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
	var banHistory []*models.BanHistory
	if err := q.selectAll(banHistory, "ban_history", "start_time", "user_id=$1", userID); err != nil {
		return nil, err
	}

	return banHistory, nil
}

func (q *PostgresQueries) GetCurrentBanFromUserID(userID int64) (*models.BanHistory, error) {
	return nil, nil // TODO
}

func (q *PostgresQueries) DeleteBanHistory(startTime time.Time, userID int64) error { // TODO: instead of startTime, use n where this ban is the n-th time?
	if err := q.deleteWhere("ban_history", true, "start_time=$1 AND user_id=$2", startTime, userID); err != nil {
		return err // TODO: error handling when ban does not exist
	}

	return nil
}
