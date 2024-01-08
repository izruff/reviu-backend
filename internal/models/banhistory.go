package models

import (
	"database/sql"
	"time"
)

type BanHistory struct {
	StartTime   time.Time      `db:"start_time"`
	EndTime     time.Time      `db:"end_time"`
	Reason      sql.NullString `db:"reason"`
	UserID      int32          `db:"user_id"`
	ModeratorID int32          `db:"created_at"`
}
