package models

import "time"

type Subscription struct {
	TopicID   int32     `db:"topic_id"`
	UserID    int32     `db:"user_id"`
	CreatedAt time.Time `db:"created_at"`
}
