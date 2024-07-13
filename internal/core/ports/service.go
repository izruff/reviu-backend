package ports

import (
	"time"

	"github.com/izruff/reviu-backend/internal/core/domain"
	"gopkg.in/guregu/null.v3"
)

type Service interface {
	Login(usernameOrEmail string, password string) (int64, string, *SvcError)
	Signup(email string, username string, password string) (int64, string, *SvcError)

	GetUserByID(id int64) (*domain.User, *SvcError)
	GetUserIDByUsername(username string) (int64, *SvcError)
	UpdateUserByID(id int64, updatedUser *domain.User) *SvcError
	BanUserByID(id int64, moderatorID int64, reason string, startTime time.Time, endTime time.Time) *SvcError
	SearchUsers(options *domain.SearchUsersOptions) ([]domain.User, *SvcError)
	FollowUserByID(followerID int64, followingID int64) *SvcError
	UnfollowUserByID(followerID int64, followingID int64) *SvcError
	GetUserFollowers(id int64) ([]domain.User, *SvcError)
	GetUserFollowerCount(id int64) (int64, *SvcError)
	GetUserFollowings(id int64) ([]domain.User, *SvcError)
	GetUserFollowingCount(id int64) (int64, *SvcError)
	GetUserPostCount(id int64) (int64, *SvcError)
	GetUserRating(id int64) (int64, *SvcError)
	GetUserSubscriptions(id int64) ([]domain.Subscription, *SvcError)
	GetUserBookmarks(id int64) ([]domain.Bookmark, *SvcError)

	CreateTopic(topic string, hub string) (int64, *SvcError)
	GetTopicByID(id int64) (*domain.Topic, *SvcError)
	UpdateTopicByID(id int64, description string) *SvcError
	SearchTopics(options *domain.SearchTopicsOptions) ([]domain.Topic, *SvcError)

	CreatePost(title string, content string, authorID int64, topic string, hub string, tags []string) (int64, *SvcError)
	GetPostByID(id int64) (*domain.Post, *SvcError)
	GetPostInteractionsByUserID(id int64, userID int64) (bool, *null.Bool, *SvcError)
	UpdatePostByID(id int64, updatedPost *domain.Post) *SvcError
	MarkPostAsDeletedByID(id int64, reasonForDeletion string, moderatorID int64) *SvcError
	ViewPost(id int64, userID int64) *SvcError
	VotePost(id int64, userID int64, up null.Bool) *SvcError
	SearchPosts(options *domain.SearchPostsOptions) ([]domain.Post, *SvcError)
	GetPostsByAuthorID(authorID int64) *SvcError
	BookmarkPostWithID(postID int64, userID int64) *SvcError

	CreateComment(content string, authorID int64, postID null.Int, parentCommentID null.Int) (int64, *SvcError)
	GetCommentByID(commentID int64) (*domain.Comment, *SvcError)
	UpdateCommentByID(commentID int64, content string) *SvcError
	MarkCommentAsDeletedByID(commentID int64, postID int64, reasonForDeletion string, moderatorID int64) *SvcError
	VoteComment(id int64, userID int64, up null.Bool) *SvcError
	SearchComments(options *domain.SearchCommentsOptions) ([]domain.Comment, *SvcError)

	CreateTag(tag string, hub string) (int64, *SvcError)
	GetTagByID(id int64) (*domain.Tag, *SvcError)
	SearchTags(options *domain.SearchTagsOptions) ([]domain.Tag, *SvcError)
}

type SvcError struct {
	Code    int
	Message string
}
