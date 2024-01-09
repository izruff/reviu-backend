package models

import "time"

type Subscription struct {
	TopicID   int32     `db:"topic_id" json:"topicId"`
	UserID    int32     `db:"user_id" json:"userId"`
	CreatedAt time.Time `db:"created_at" json:"createdAt"`
}
