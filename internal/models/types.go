package models

// TODOS

type SearchUsersOptions struct {
	Query string `form:"q"`

	// similarity (default): sort by the similarity of word composition, using pg_trgm extension.
	// popularity: sort roughly by user followings and total vote counts.
	SortBy string `form:"sort-by"`

	// match-left: only shows results which matches the left part of the string and passes the no-match criteria.
	// match-substring: only shows results which matches some substring and passes the no-match criteria.
	// no-match (default): shows results which passes a certain threshold of similarness determined by pg_trgm.
	MustMatch string `form:"must-match"`
}

type SearchPostsOptions struct {
	Query string `form:"q"`

	// similarity: sort by the similarity of word composition, using pg_trgm extension.
	// popularity: sort roughly by view, comment, or vote counts.
	// age-asc (default): sort from the newest post.
	// age-desc: sort from the oldest post.
	SortBy string `form:"sort-by"`

	// title: compare query with only title.
	// all: compare query with both title and content.
	MatchWith string `form:"match-with"`

	Topics []string `form:"topics"`
	Tags   []string `form:"tags"`

	// deleted posts are not shown
}

type SearchCommentsOptions struct {
	PostID int64  // will be included from context.Param() function.
	Query  string `form:"q"`

	// similarity: sort by the similarity of word composition, using pg_trgm extension.
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

type SearchTagsOptions struct { // TODO
	Query string `form:"q"`

	// similarity (default): sort by the similarity of word composition, using pg_trgm extension.
	// popularity: sort  by number of posts using it.
	SortBy string `form:"sort-by"`

	// left: only shows results which matches the left part of the string and passes the no-match criteria.
	// substring: only shows results which matches some substring and passes the no-match criteria.
	// none (default): shows results which passes a certain threshold of similarness determined by pg_trgm.
	MustMatch string `form:"must-match"`
}
