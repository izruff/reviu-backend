package models

// TODOS

type SearchUsersOptions struct {
	Query      string `form:"q"`
	SortBy     string `form:"sort"`
	ExactMatch string `form:"exact"`
}

type SearchPostsOptions struct { // TODO
	Query      string `form:"q"`
	SortBy     string `form:"sort"`
	ExactMatch string `form:"exact"`
}

type SearchCommentsOptions struct { // TODO
	PostID     int64
	Query      string `form:"q"`
	SortBy     string `form:"sort"`
	ExactMatch string `form:"exact"`
}

type SearchTopicsOptions struct { // TODO
	Query      string `form:"q"`
	SortBy     string `form:"sort"`
	ExactMatch string `form:"exact"`
}

type SearchTagsOptions struct { // TODO
	Query      string `form:"q"`
	SortBy     string `form:"sort"`
	ExactMatch string `form:"exact"`
}
