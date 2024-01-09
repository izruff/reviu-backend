package models

import "time"

type Relation struct {
	FollowerID  int32     `db:"follower_id" json:"followerId"`
	FollowingID int32     `db:"following_id" json:"followingId"`
	CreatedAt   time.Time `db:"created_at" json:"createdAt"`
}
