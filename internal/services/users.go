package services

import (
	"regexp"

	"github.com/izruff/reviu-backend/internal/models"
	"github.com/izruff/reviu-backend/internal/utils"
)

func isValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`)
	return emailRegex.MatchString(email)
}

func (s *APIServices) Login(usernameOrEmail string, password string) (int32, string, *SvcError) {
	var userID int32
	var err error
	switch isValidEmail(usernameOrEmail) {
	case true:
		userID, err = s.queries.GetUserIDByEmail(usernameOrEmail)
	case false:
		userID, err = s.queries.GetUserIDByUsername(usernameOrEmail)
	}
	if err != nil {
		return 0, "", newErrInternal(err) // TODO: error handling when user does not exist
	}

	user, err := s.queries.GetUserByID(userID) // may decide to return the whole user at some point
	if err != nil {
		return 0, "", newErrInternal(err)
	}

	if err := utils.AssertValidPassword(password, user.PasswordHash.String); err != nil {
		return 0, "", newErrInvalidCredentials("invalid password")
	}

	// TODO: generate token
	token := ""

	return userID, token, nil
}

func (s *APIServices) Signup(email string, username string, password string) (int32, string, *SvcError) {
	passwordHash, err := utils.GetPasswordHash(password)
	if err != nil {
		return 0, "", newErrInternal(err)
	}

	newUser := &models.User{
		Email:        *NewString(email),
		Username:     *NewString(username),
		PasswordHash: *NewString(passwordHash),
		ModRole:      *NewBool(false),
	}

	userID, err := s.queries.CreateUser(newUser)
	// TODO: error handling when user already exists
	if err != nil {
		return 0, "", newErrInternal(err)
	}

	// TODO: generate token
	token := ""

	return userID, token, nil
}
