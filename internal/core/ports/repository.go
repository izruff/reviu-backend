package ports

import (
	"time"

	"github.com/izruff/reviu-backend/internal/core/domain"
	"gopkg.in/guregu/null.v3"
)

type Repository interface {
	CreateUser(newUser *domain.User) (int64, error)
	GetUserByID(id int64) (*domain.User, error)
	GetUserIDByEmail(email string) (int64, error)
	GetUserIDByUsername(username string) (int64, error)
	GetUsersWithOptions(options *domain.SearchUsersOptions) ([]domain.User, error)
	UpdateUserByID(updatedUser *domain.User) error
	DeleteUserByID(id int64) error

	CreateTopic(newTopic *domain.Topic) (int64, error)
	GetTopicByID(id int64) (*domain.Topic, error)
	GetTopicID(topic string, hub string) (int64, error)
	GetTopicsWithOptions(options *domain.SearchTopicsOptions) ([]domain.Topic, error)
	UpdateTopicByID(id int64, description string) error

	CreatePost(newPost *domain.Post) (int64, error)
	GetPostByID(id int64) (*domain.Post, error)
	GetPostsWithOptions(options *domain.SearchPostsOptions) ([]domain.Post, error)
	CountPostsFromAuthorID(userID int64) (int64, error)
	UpdatePostByID(updatedPost *domain.Post) error
	MarkPostAsDeletedByID(id int64, reason string, moderatorID int64) error

	CreateComment(newComment *domain.Comment) (int64, error)
	GetCommentByID(id int64) (*domain.Comment, error)
	GetCommentsWithOptions(options *domain.SearchCommentsOptions) ([]domain.Comment, error)
	UpdateCommentByID(updatedComment *domain.Comment) error
	MarkCommentAsDeletedByID(id int64, reason string, moderatorID int64) error

	CreateTag(newTag *domain.Tag) (int64, error)
	GetTagByID(id int64) (*domain.Tag, error)
	GetTagsWithOptions(options *domain.SearchTagsOptions) ([]domain.Tag, error)

	CreateTaggedPost(newTaggedPost *domain.TaggedPost) error
	GetTagsFromPostID(postID int64) ([]domain.Tag, error)
	DeleteTagFromPost(postID int64, tagID int64) error

	CreatePostView(newView *domain.PostView) error
	GetPostViewValue(postID int64, userID int64) (bool, error)
	GetViewsFromPostID(postID int64) ([]domain.PostView, error)
	CountViewsFromPostID(postID int64) (int64, error)
	DeletePostView(postID int64, userID int64) error

	CreatePostVote(newVote *domain.PostVote) error
	GetPostVoteValue(postID int64, userID int64) (*null.Bool, error)
	GetVotesFromPostID(postID int64) ([]domain.PostVote, error)
	CountVotesFromPostID(postID int64) (int64, int64, error)
	UpdatePostVote(up bool, postID int64, userID int64) error
	DeletePostVote(postID int64, userID int64) error

	CreateCommentVote(newVote *domain.CommentVote) error
	GetCommentVoteValue(commentID int64, userID int64) (*null.Bool, error)
	GetVotesFromCommentID(commentID int64) ([]domain.CommentVote, error)
	CountVotesFromCommentID(commentID int64) (int64, int64, error)
	UpdateCommentVote(up bool, commentID int64, userID int64) error
	DeleteCommentVote(commentID int64, userID int64) error

	CreateRelation(newRelation *domain.Relation) error
	GetFollowersFromUserID(userID int64) ([]domain.Relation, error)
	CountFollowersFromUserID(userID int64) (int64, error)
	GetFollowingsFromUserID(userID int64) ([]domain.Relation, error)
	CountFollowingsFromUserID(userID int64) (int64, error)
	DeleteRelation(followerID int64, followingID int64) error

	CreateSubscription(newSubscription *domain.Subscription) error
	GetSubscribersFromTopicID(topicID int64) ([]domain.Subscription, error)
	CountSubscribersFromTopicID(topicID int64) (int64, error)
	GetSubscribedTopicsFromUserID(userID int64) ([]domain.Subscription, error)
	CountSubscribedTopicsFromUserID(userID int64) (int64, error)
	DeleteSubscription(topicID int64, userID int64) error

	CreateBookmark(newBookmark *domain.Bookmark) error
	GetBookmarksFromUserID(userID int64) ([]domain.Bookmark, error)
	CountBookmarksFromUserID(userID int64) (int64, error)
	DeleteBookmark(postID int64, userID int64) error

	CreateBanHistory(newBanHistory *domain.BanHistory) error
	GetBanHistoryFromUserID(userID int64) ([]domain.BanHistory, error)
	GetCurrentBanFromUserID(userID int64) (*domain.BanHistory, error)
	DeleteBanHistory(startTime time.Time, userID int64) error
}
