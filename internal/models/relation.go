package models

import "time"

type Relation struct {
	FollowerID  int32     `db:"follower_id"`
	FollowingID int32     `db:"following_id"`
	CreatedAt   time.Time `db:"created_at"`
}
