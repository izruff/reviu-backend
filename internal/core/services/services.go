package services

import (
	"net/http"
	"strings"

	"github.com/izruff/reviu-backend/internal/core/ports"
)

type APIServices struct {
	repo ports.Repository
}

func NewAPIServices(repo ports.Repository) *APIServices {
	return &APIServices{
		repo: repo,
	}
}

type SvcError struct {
	Code    int
	Message string
}

// Internal server error; for any unexpected error that is not categorized here
func newErrInternal(err error) *SvcError {
	return &SvcError{
		Code:    http.StatusInternalServerError,
		Message: err.Error(),
	}
}

// Error for invalid, expired, or non-existent credentials
func newErrInvalidCredentials(errMessage string) *SvcError {
	return &SvcError{
		Code:    http.StatusUnauthorized,
		Message: errMessage,
	}
}

// Error for invalid user input
// Message format: "invalid user input: email, username"
func newErrInvalidUserInput(invalidFields []string) *SvcError {
	return &SvcError{
		Code:    http.StatusBadRequest,
		Message: "invalid user input: " + strings.Join(invalidFields, ", "),
	}
}
