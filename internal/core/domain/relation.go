package models

import "gopkg.in/guregu/null.v3"

type Relation struct {
	FollowerID  null.Int  `db:"follower_id" json:"followerId"`
	FollowingID null.Int  `db:"following_id" json:"followingId"`
	CreatedAt   null.Time `db:"created_at" json:"createdAt"`
}
