package domain

import "gopkg.in/guregu/null.v3"

type SearchUsersOptions struct {
	Query string `form:"q"`

	// similarity (default): sort by the similarity of word composition, using pg_trgm extension.
	// popularity: sort roughly by user followings and total vote counts.
	SortBy string `form:"sort-by"`

	// left: only shows results which matches the left part of the string and passes the no-match criteria.
	// substring: only shows results which matches some substring and passes the no-match criteria.
	// none (default): shows results which passes a certain threshold of similarness determined by pg_trgm.
	MustMatch string `form:"must-match"`
}

type SearchPostsOptions struct {
	Query string `form:"q"`

	// similarity: sort by the similarity of word composition, using pg_trgm extension.
	// popularity: sort roughly by time created, view, vote, and comment count.
	// age-asc (default): sort from the newest post.
	// age-desc: sort from the oldest post.
	SortBy string `form:"sort-by"`

	// title: compare query with only title.
	// all (default): compare query with both title and content.
	MatchWith string `form:"match-with"`

	Authors []string `form:"authors"`
	Topics  []string `form:"topics"`
	Tags    []string `form:"tags"`

	// deleted posts are not shown
}

type SearchCommentsOptions struct {
	PostID          null.Int
	ParentCommentID null.Int
	Query           string `form:"q"`

	// similarity: sort by the similarity of word composition, using pg_trgm extension.
	// popularity: sort roughly by time created, vote, and comment count.
	// age-asc (default): sort from the newest post.
	// age-desc: sort from the oldest post.
	SortBy string `form:"sort-by"`

	// deleted comments are not shown
}

type SearchTopicsOptions struct {
	Query string `form:"q"`

	// similarity (default): sort by the similarity of word composition, using pg_trgm extension.
	// popularity: sort roughly by subscription and post counts.
	SortBy string `form:"sort-by"`

	// left: only shows results which matches the left part of the string and passes the no-match criteria.
	// substring: only shows results which matches some substring and passes the no-match criteria.
	// none (default): shows results which passes a certain threshold of similarness determined by pg_trgm.
	MustMatch string `form:"must-match"`
}

type SearchTagsOptions struct {
	Query string `form:"q"`

	// similarity (default): sort by the similarity of word composition, using pg_trgm extension.
	// popularity: sort  by number of posts using it.
	SortBy string `form:"sort-by"`

	// left: only shows results which matches the left part of the string and passes the no-match criteria.
	// substring: only shows results which matches some substring and passes the no-match criteria.
	// none (default): shows results which passes a certain threshold of similarness determined by pg_trgm.
	MustMatch string `form:"must-match"`
}
