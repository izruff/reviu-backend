package services

import "net/http"

type SvcError struct {
	Code    int
	Message string
}

// TODO: create more errors
func newErrInternal(err error) *SvcError {
	return &SvcError{
		Code:    http.StatusInternalServerError,
		Message: err.Error(),
	}
}

func newErrInvalidCredentials(errMessage string) *SvcError {
	return &SvcError{
		Code:    http.StatusUnauthorized,
		Message: errMessage,
	}
}

func newErrInvalidUserInput(errMessage string) *SvcError {
	return &SvcError{
		Code:    http.StatusBadRequest,
		Message: errMessage,
	}
}
