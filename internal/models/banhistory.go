package models

import "gopkg.in/guregu/null.v3"

type BanHistory struct {
	StartTime   null.Time   `db:"start_time" json:"startTime"`
	EndTime     null.Time   `db:"end_time" json:"endTime"`
	Reason      null.String `db:"reason" json:"reason"`
	UserID      null.Int    `db:"user_id" json:"userId"`
	ModeratorID null.Int    `db:"created_at" json:"createdAt"`
}
