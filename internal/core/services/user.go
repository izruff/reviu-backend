package services

import (
	"regexp"
	"time"

	"github.com/izruff/reviu-backend/internal/core/domain"
	"github.com/izruff/reviu-backend/internal/utils"
	"gopkg.in/guregu/null.v3"
)

func isValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`)
	return emailRegex.MatchString(email)
}

func (s *APIServices) Login(usernameOrEmail string, password string) (int64, string, *SvcError) {
	var userID int64
	var err error
	switch isValidEmail(usernameOrEmail) {
	case true:
		userID, err = s.repo.GetUserIDByEmail(usernameOrEmail)
	case false:
		userID, err = s.repo.GetUserIDByUsername(usernameOrEmail)
	}
	if err != nil {
		return 0, "", newErrInternal(err) // TODO: error handling when user does not exist
	}

	user, err := s.repo.GetUserByID(userID) // may decide to return the whole user at some point
	if err != nil {
		return 0, "", newErrInternal(err)
	}

	if err := utils.AssertValidPassword(password, user.PasswordHash.String); err != nil {
		return 0, "", newErrInvalidCredentials("invalid password")
	}

	token, err := utils.GenerateJWT(userID)
	if err != nil {
		return 0, "", newErrInternal(err)
	}

	return userID, token, nil
}

func (s *APIServices) Signup(email string, username string, password string) (int64, string, *SvcError) {
	passwordHash, err := utils.GetPasswordHash(password)
	if err != nil {
		return 0, "", newErrInternal(err)
	}

	newUser := &domain.User{
		Email:        null.NewString(email, true),
		Username:     null.NewString(username, true),
		PasswordHash: null.NewString(passwordHash, true),
		ModRole:      null.NewBool(false, true),
	}

	userID, err := s.repo.CreateUser(newUser)
	// TODO: error handling when user already exists
	if err != nil {
		return 0, "", newErrInternal(err)
	}

	token, err := utils.GenerateJWT(userID)
	if err != nil {
		return 0, "", newErrInternal(err)
	}

	return userID, token, nil
}

func (s *APIServices) GetUserByID(id int64) (*domain.User, *SvcError) {
	user, err := s.repo.GetUserByID(id)
	if err != nil {
		// TODO: error handling when user does not exist
		return nil, newErrInternal(err)
	}

	return user, nil
}

func (s *APIServices) GetUserIDByUsername(username string) (int64, *SvcError) {
	userID, err := s.repo.GetUserIDByUsername(username)
	if err != nil {
		return 0, newErrInternal(err) // TODO: error handling when user does not exist
	}

	return userID, nil
}

func (s *APIServices) UpdateUserByID(id int64, updatedUser *domain.User) *SvcError {
	// TODO: error handling when there are no changes
	updatedUser.ID.Int64 = id
	updatedUser.ID.Valid = true
	if err := s.repo.UpdateUserByID(updatedUser); err != nil {
		// TODO: error handling when user does not exist
		return newErrInternal(err)
	}

	return nil
}

func (s *APIServices) BanUserByID(id int64, moderatorID int64, reason string, startTime time.Time, endTime time.Time) *SvcError {
	return nil // TODO
}

func (s *APIServices) SearchUsers(options *domain.SearchUsersOptions) ([]domain.User, *SvcError) {
	users, err := s.repo.GetUsersWithOptions(options)
	if err != nil {
		return nil, newErrInternal(err) // TODO: error handling when there are incorrect options
	}

	return users, nil
}

func (s *APIServices) FollowUserByID(followerID int64, followingID int64) *SvcError {
	newRelation := &domain.Relation{
		FollowerID:  null.NewInt(followerID, true),
		FollowingID: null.NewInt(followingID, true),
	}

	if err := s.repo.CreateRelation(newRelation); err != nil {
		// TODO: error handling when user does not exist
		return newErrInternal(err)
	}

	return nil
}

func (s *APIServices) UnfollowUserByID(followerID int64, followingID int64) *SvcError {
	if err := s.repo.DeleteRelation(followerID, followingID); err != nil {
		// TODO: error handling when relation does not exist
		return newErrInternal(err)
	}

	return nil
}

func (s *APIServices) GetUserFollowers(id int64) ([]domain.User, *SvcError) {
	followers, err := s.repo.GetFollowersFromUserID(id)
	if err != nil {
		return nil, newErrInternal(err) // TODO: error handling when user does not exist
	}

	users := []domain.User{}
	for _, relation := range followers {
		user, err := s.GetUserByID(relation.FollowerID.Int64)
		if err != nil {
			return nil, err
		}
		users = append(users, *user)
	}

	return users, nil
}

func (s *APIServices) GetUserFollowerCount(id int64) (int64, *SvcError) {
	count, err := s.repo.CountFollowersFromUserID(id)
	if err != nil {
		return 0, newErrInternal(err) // TODO: error handling when user does not exist
	}

	return count, nil
}

func (s *APIServices) GetUserFollowings(id int64) ([]domain.User, *SvcError) {
	followings, err := s.repo.GetFollowingsFromUserID(id)
	if err != nil {
		return nil, newErrInternal(err) // TODO: error handling when user does not exist
	}

	users := []domain.User{}
	for _, relation := range followings {
		user, err := s.GetUserByID(relation.FollowingID.Int64)
		if err != nil {
			return nil, err
		}
		users = append(users, *user)
	}

	return users, nil
}

func (s *APIServices) GetUserFollowingCount(id int64) (int64, *SvcError) {
	count, err := s.repo.CountFollowingsFromUserID(id)
	if err != nil {
		return 0, newErrInternal(err) // TODO: error handling when user does not exist
	}

	return count, nil
}

func (s *APIServices) GetUserPostCount(id int64) (int64, *SvcError) {
	count, err := s.repo.CountPostsFromAuthorID(id)
	if err != nil {
		return 0, newErrInternal(err) // TODO: error handling when user does not exist
	}

	return count, nil
}

func (s *APIServices) GetUserRating(id int64) (int64, *SvcError) {
	// TODO: create the column in database

	return 0, nil
}

func (s *APIServices) GetUserSubscriptions(id int64) ([]domain.Subscription, *SvcError) {
	subscriptions, err := s.repo.GetSubscribedTopicsFromUserID(id)
	if err != nil {
		return nil, newErrInternal(err) // TODO: error handling when user does not exist
	}

	return subscriptions, nil
}

func (s *APIServices) GetUserBookmarks(id int64) ([]domain.Bookmark, *SvcError) {
	bookmarks, err := s.repo.GetBookmarksFromUserID(id)
	if err != nil {
		return nil, newErrInternal(err) // TODO: error handling when user does not exist
	}

	return bookmarks, nil
}
