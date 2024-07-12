package domain

import "gopkg.in/guregu/null.v3"

type Subscription struct {
	TopicID   null.Int  `db:"topic_id" json:"topicId"`
	UserID    null.Int  `db:"user_id" json:"userId"`
	CreatedAt null.Time `db:"created_at" json:"createdAt"`
}
