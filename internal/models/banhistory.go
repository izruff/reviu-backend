package models

import (
	"database/sql"
	"time"
)

type BanHistory struct {
	StartTime   time.Time      `db:"start_time" json:"startTime"`
	EndTime     time.Time      `db:"end_time" json:"endTime"`
	Reason      sql.NullString `db:"reason" json:"reason"`
	UserID      int32          `db:"user_id" json:"userId"`
	ModeratorID int32          `db:"created_at" json:"createdAt"`
}
