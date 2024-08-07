package http

import "gopkg.in/guregu/null.v3"

// For account handlers

type loginJSON struct {
	UsernameOrEmail string `json:"usernameOrEmail" binding:"required"`
	Password        string `json:"password" binding:"required"`
}

type signupJSON struct {
	Email    string `json:"email" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// For authorized handlers

type updateUserProfileJSON struct {
	Nickname null.String `json:"nickname"`
	About    null.String `json:"about"`
}

type followOrUnfollowUserJSON struct {
	FollowingID int64 `json:"followingId" binding:"required"`
}

type createPostJSON struct {
	Title   string   `json:"title" binding:"required"`
	Content string   `json:"content" binding:"required"`
	Topic   string   `json:"topic" binding:"required"`
	Hub     string   `json:"hub" binding:"required"`
	Tags    []string `json:"tags"`
}

type editPostJSON struct {
	PostID        int64       `json:"postId" binding:"required"`
	Title         null.String `json:"title"`
	Content       null.String `json:"content"`
	AddedTagsID   []int64     `json:"addedTagsId"`
	RemovedTagsID []int64     `json:"removedTagsId"`
}

type replyToPostJSON struct {
	Content string `json:"content" binding:"required"`
	PostID  int64  `json:"postId" binding:"required"`
}

type votePostJSON struct {
	Up     null.Bool `json:"up"`
	PostID int64     `json:"postId" binding:"required"`
}

type bookmarkPostJSON struct {
	PostID int64 `json:"postId" binding:"required"`
}

type replyToCommentJSON struct {
	Content         string `json:"content" binding:"required"`
	ParentCommentID int64  `json:"parentCommentID" binding:"required"`
}

type voteCommentJSON struct {
	Up        null.Bool `json:"up"`
	CommentID int64     `json:"commentId" binding:"required"`
}

type editCommentJSON struct {
	CommentID int64  `json:"commentId" binding:"required"`
	Content   string `json:"content" binding:"required"`
}

type createTopicJSON struct {
	Topic string `json:"topic" binding:"required"`
	Hub   string `json:"hub" binding:"required"`
}

type postInteractionsResponse struct {
	Viewed bool       `json:"viewed"`
	Voted  *null.Bool `json:"voted"`
}
