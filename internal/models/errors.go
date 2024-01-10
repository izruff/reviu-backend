package models

import "fmt"

type ErrInvalidValue struct {
	property string
	message  string
}

func NewErrInvalidValue(property, message string) error {
	return &ErrInvalidValue{
		property: property,
		message:  message,
	}
}

func (e *ErrInvalidValue) Error() string {
	return fmt.Sprintf("invalid value for property %s: %s", e.property, e.message)
}
