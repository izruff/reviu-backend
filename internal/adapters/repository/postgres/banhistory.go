package postgres

import (
	"time"

	"github.com/izruff/reviu-backend/internal/core/domain"
)

func (r *PostgresRepository) CreateBanHistory(newBanHistory *domain.BanHistory) error {
	// TODO: error handling when form is incomplete or already exists
	if _, err := r.create("ban_history", []string{"start_time", "end_time", "reason", "user_id", "moderator_id"}, false, newBanHistory); err != nil {
		return err
	}

	return nil
}

func (r *PostgresRepository) GetBanHistoryFromUserID(userID int64) ([]domain.BanHistory, error) {
	banHistory := []domain.BanHistory{}
	if err := r.selectAll(&banHistory, "ban_history", "start_time", "user_id=$1", "", userID); err != nil {
		return nil, err
	}

	return banHistory, nil
}

func (r *PostgresRepository) GetCurrentBanFromUserID(userID int64) (*domain.BanHistory, error) {
	return nil, nil // TODO
}

func (r *PostgresRepository) DeleteBanHistory(startTime time.Time, userID int64) error { // TODO: instead of startTime, use n where this ban is the n-th time?
	if err := r.deleteWhere("ban_history", true, "start_time=$1 AND user_id=$2", startTime, userID); err != nil {
		return err // TODO: error handling when ban does not exist
	}

	return nil
}
